package main

import "fmt"

func SliceNumber(number []int) {
	var newNumber int

	for i := 0; i <= len(number); i++ {
		newNumber += i
	}
	fmt.Println(newNumber)
}

func main() {
	caseSatu := []int{1, 2, 3, 4, 5}
	caseDua := []int{1, 2, 3}
	caseTiga := []int{}
	SliceNumber(caseSatu)
	SliceNumber(caseDua)
	SliceNumber(caseTiga)
}
