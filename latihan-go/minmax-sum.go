package main

import "fmt"

func main() {
	array := [...]int{10, 120, 14, 50, 5}

	min := array[0]
	max := array[0]

	for _, value := range array {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}
	fmt.Println((max + min) + (max - min))
}
