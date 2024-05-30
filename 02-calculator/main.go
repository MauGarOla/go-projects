package main

import (
	"errors"
	"fmt"
)

func Calculator(x, y float32, operation string) (float32, error) {
	var result float32

	switch operation {
	case "sum":
		result = x + y
	case "rest":
		result = x - y
	case "multiplication":
		result = x * y
	case "division":
		if y == 0 {
			return 0, errors.New("math error")
		}
		result = x / y
	default:
		return 0, errors.New("parameters error")
	}
	return result, nil
}

func main() {

	var x, y float32
	var operation string

	fmt.Print("First number: ")
	fmt.Scan(&x)

	fmt.Print("Second number: ")
	fmt.Scan(&y)

	fmt.Print("Operation (Sum, rest, multiplication, division): ")
	fmt.Scan(&operation)

	result, err := Calculator(x, y, operation)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Your result is %v", result)
	}
}
