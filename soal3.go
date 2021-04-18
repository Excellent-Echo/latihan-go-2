package main

import "fmt"

func main() {
	arr := [6]int{1, 205, 5000, 41, 200, 2000}
	var max, min, addition, subtraction, result int
	max = arr[0]
	min = arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}
	addition = max + min
	subtraction = max - min
	result = addition + subtraction
	fmt.Printf("penjumlahan nilai max dan min adalah %v + %v = %v\n", max, min, addition)
	fmt.Printf("pengurangan nilai max dan min adalah %v + %v = %v\n", max, min, subtraction)
	fmt.Printf("%v + %v = %v\n", addition, subtraction, result)
}
