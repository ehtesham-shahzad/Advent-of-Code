package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	var firstNumber, lastNumber uint8
	var firstNumberSet bool
	var result uint16

	for i := 0; i < len(data); i++ {

		if data[i] >= 48 && data[i] <= 57 {
			if !firstNumberSet {
				firstNumber = uint8(data[i]) - 48
				firstNumberSet = true
			}
			lastNumber = uint8(data[i]) - 48
		}

		if string(data[i]) == "\n" || i == len(data)-1 {
			firstNumber = firstNumber * 10
			result = result + uint16(firstNumber) + uint16(lastNumber)
			firstNumber, lastNumber = 0, 0
			firstNumberSet = false
		}
	}
	fmt.Printf("Result: %d", result)
}
