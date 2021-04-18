package main

import (
	"fmt"
)

func main() {
	var number = [...]int{10, 120, 14, 50, 5}

	minNumber := number[0]
	maxNumber := number[0]

	for _, element := range number {
		if element < minNumber {
			minNumber = element
		}

		if element > maxNumber {
			maxNumber = element
		}
	}

	fmt.Println("Isi dari index array terkecil =", minNumber)
	fmt.Println("Isi dari index array terbesar =", maxNumber)

	addition := maxNumber + minNumber
	substraction := maxNumber - minNumber
	result := addition + substraction

	fmt.Println(maxNumber, "+", minNumber, "=", addition)
	fmt.Println(maxNumber, "-", minNumber, "=", substraction)
	fmt.Println("Hasil =", result)
}