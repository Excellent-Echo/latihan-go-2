package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	var result int

	for _, num := range arr {
		result += num
	}
	fmt.Println(result)
}
