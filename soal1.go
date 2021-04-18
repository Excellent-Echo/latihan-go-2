package main

import "fmt"

func main() {
	var (
		angka = [5]int{1, 2, 3, 4, 5}
		hasil int
	)
	for i := 0; i <= len(angka); i++ {
		hasil += len(angka)

	}
	fmt.Println(hasil)
}
