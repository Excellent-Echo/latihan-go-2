package main

import (
	"fmt"
)

func main() {
	var year int

	fmt.Print("input year: ")
	fmt.Scanf("%d", &year)

	if year <= 0 {
		fmt.Println("Tahun invalid")
	} else if year < 100 {
		fmt.Printf("Tahun %d termasuk Abad ke-1", year)
	} else if year%100 == 0 {
		century := year / 100
		fmt.Printf("Tahun %d termasuk abad ke %d", year, century)
	} else {
		century := (year / 100) + 1
		fmt.Printf("Tahun %d termasuk abad ke %d", year, century)
	}
}
