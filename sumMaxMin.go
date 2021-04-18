package main

import "fmt"

func main() {
	array := []int{10, 120, 14, 50, 5}
	max := array[0]
	min := array[0]

	for _, values := range array {
		if values < min {
			min = values
		}
		if values > max {
			max = values
		}
	}
	fmt.Print((max + min) + (max - min))
}
