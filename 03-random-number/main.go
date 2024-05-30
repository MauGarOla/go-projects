package main

import (
	"fmt"
	"math/rand"
)

func Random(numbers ...float32) float32 {

	var min, max float32

	switch len(numbers) {
	case 0:
		min, max = 0, 100
	case 1:
		min, max = 1, numbers[0]
	case 2:
		min, max = numbers[0], numbers[1]
	default:
		panic("to many parameters")
	}

	if min > max {
		min, max = max, min
	}

	return min + rand.Float32()*(max-min)
}

func main() {

	var min, max float32

	fmt.Print("Min: ")
	fmt.Scan(&min)

	fmt.Print("Max: ")
	fmt.Scan(&max)

	fmt.Print(Random(min, max))
}
