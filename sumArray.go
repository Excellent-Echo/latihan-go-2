package main

import (
	"fmt"
)

func main() {
	array := []int{3, 3, 8, 3, 8, 2}
	hasilPenjumlahan := 0

	for _, num := range array {
		hasilPenjumlahan += num
	}
	fmt.Println(hasilPenjumlahan) //27
}