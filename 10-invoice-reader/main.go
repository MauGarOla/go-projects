package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	pdf "github.com/ledongthuc/pdf"
	"github.com/sqweek/dialog"
)

type InvoiceItem struct {
	InvoiceNumber string
	IssueDate     string
	Vendor        string
	Code          string
	Description   string
	Quantity      string
	UnitPrice     string
	TotalPrice    string
}

func main() {
	// Obtener la ruta del archivo PDF
	pdfPath, err := getFilePath("Select PDF file")
	if err != nil {
		log.Fatalf("Failed to open PDF file dialog: %v", err)
	}

	// Obtener la ruta del archivo CSV
	csvPath, err := getFilePath("Select where to save CSV file")
	if err != nil {
		log.Fatalf("Failed to open CSV file dialog: %v", err)
	}

	// Extract text from PDF
	extractedText, err := extractTextFromPDF(pdfPath)
	if err != nil {
		log.Fatalf("Failed to extract text from PDF: %v", err)
	}

	// Parse extracted text into invoice items
	invoiceItems := parseInvoiceText(extractedText)

	// Save invoice items to CSV
	err = saveToCSV(invoiceItems, csvPath)
	if err != nil {
		log.Fatalf("Failed to save to CSV: %v", err)
	}

	fmt.Println("Invoice data successfully saved to CSV.")
}

func getFilePath(title string) (string, error) {
	filePath, err := dialog.File().Title(title).Load()
	if err != nil {
		return "", err
	}
	fmt.Printf("Selected file: %s\n", filePath)
	return filePath, nil
}

func extractTextFromPDF(pdfPath string) (string, error) {
	f, r, err := pdf.Open(pdfPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf strings.Builder
	totalPage := r.NumPage()
	for i := 1; i <= totalPage; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}
		text, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		buf.WriteString(text)
	}

	return buf.String(), nil
}

func parseInvoiceText(text string) []InvoiceItem {
	lines := strings.Split(text, "\n")
	var invoiceItems []InvoiceItem
	var currentInvoice InvoiceItem

	for _, line := range lines {
		if strings.Contains(line, "Invoice number:") {
			currentInvoice.InvoiceNumber = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "am") {
			currentInvoice.IssueDate = strings.TrimSpace(strings.Split(line, "am")[1])
		} else if strings.Contains(line, "Vendor:") {
			currentInvoice.Vendor = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "Pos") {
			continue
		} else if len(strings.Fields(line)) > 4 {
			parts := strings.Fields(line)
			currentInvoice.Code = parts[1]
			currentInvoice.Description = strings.Join(parts[2:len(parts)-3], " ")
			currentInvoice.Quantity = parts[len(parts)-3]
			currentInvoice.UnitPrice = parts[len(parts)-2]
			currentInvoice.TotalPrice = parts[len(parts)-1]

			invoiceItems = append(invoiceItems, currentInvoice)
		}
	}

	return invoiceItems
}

func saveToCSV(invoiceItems []InvoiceItem, csvPath string) error {
	file, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Invoice Number", "Issue Date", "Vendor", "Code", "Description", "Quantity", "Unit Price", "Total Price"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write invoice items
	for _, item := range invoiceItems {
		record := []string{item.InvoiceNumber, item.IssueDate, item.Vendor, item.Code, item.Description, item.Quantity, item.UnitPrice, item.TotalPrice}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
