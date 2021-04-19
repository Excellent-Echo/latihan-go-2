package main

import "fmt"

func main() {
	var (
		result1, result2, result3 int
	)

	// [1, 2, 3, 4, 5] = 1 + 2 + 3 + 4 + 5 = 15
	numArr1 := []int{1, 2, 3, 4, 5}
	for _, num := range numArr1 {
		result1 += num
	}
	fmt.Println(result1)

	// [1, 2, 3] = 1 + 2 + 3 = 6
	numArr2 := []int{1, 2, 3}
	for _, num := range numArr2 {
		result2 += num
	}
	fmt.Println(result2)

	// [] = 0
	numArr3 := []int{}
	for _, num := range numArr3 {
		result3 += num
	}
	fmt.Println(result3)
}
