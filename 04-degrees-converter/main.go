package main

import "fmt"

func Converter(deg int, name string) int {
	if name == "Fahrenheit" {
		return int(float32(deg)*float32(0.5555555) + 32)
	} else if name == "Celsius" {
		return int(float32(deg-32) * float32(0.5555555))
	} else {
		return 0
	}
}

func main() {

	name := ""
	deg := 0

	fmt.Print("Do you want to have your degrees in Celsius or Fahrenheit? ")
	fmt.Scan(&name)

	fmt.Print("Wright your degrees: ")
	fmt.Scan(&deg)

	convDeg := Converter(deg, name)

	if name == "Fahrenheit" {
		fmt.Printf("%d°C is %d°F\n", deg, convDeg)
	} else if name == "Celsius" {
		fmt.Printf("%d°F is %d°C\n", deg, convDeg)
	} else {
		fmt.Println("Invalid unit, please enter 'Celsius' or 'Fahrenheit'.")
	}
}
