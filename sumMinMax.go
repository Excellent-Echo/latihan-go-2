package main

import (
	"fmt"
)

func main() {
	var number = [...]int{8, 120, 12, 20, 5}

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

	/* Isi dari index array terkecil = 5
```````Isi dari index array terbesar = 120
```````120 + 5 = 125
```````120 - 5 = 115
	   sHasil = 240 */
}