package main

import (
	"fmt"
)

func MaxMin(number []int) {
	smallesNumber := number[0]
	biggestNumber := number[0]

	for _, element := range number {
		if element < smallesNumber {
			smallesNumber = element
		}

		if element > biggestNumber {
			biggestNumber = element
		}
	}

	fmt.Println("Ini adalah array terkecil", smallesNumber)
	fmt.Println("Ini adalah array terbesar", biggestNumber)

	postiveSum := biggestNumber + smallesNumber
	negativeSum := biggestNumber - smallesNumber
	result := postiveSum + negativeSum

	fmt.Println("Positive sum :", postiveSum)
	fmt.Println("Negative sum :", negativeSum)
	fmt.Println("Hasilnya :", result)

	fmt.Println("------------------")
}

func main() {
	caseOne := []int{10, 120, 14, 50, 5}
	caseTwo := []int{0, 2, 3, 4, 5}
	MaxMin(caseOne)
	MaxMin(caseTwo)
}
