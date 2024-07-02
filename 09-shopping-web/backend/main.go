package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var Products = []Product{
	{ID: 1, Name: "Product 1", Price: 10.0},
	{ID: 2, Name: "Product 2", Price: 20.0},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func main() {
	http.HandleFunc("/api/products", getProducts)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
