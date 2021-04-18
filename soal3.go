package main

import "fmt"

func main() {
	arr := [6]int{2, 50, 3, 60, 4, 70}
	var max, min, add, sub, result int
	max = arr[0]
	min = arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] > max {
			max = arr[i]
		}

		if arr[i] < min {
			min = arr[i]
		}
	}
	add = max + min
	sub = max - min
	result = add + sub
	fmt.Printf("%v + %v = %v\n", add, sub, result)
}
