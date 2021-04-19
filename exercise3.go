package main

import "fmt"

func main() {
	var (
		numArr = [...]int{10, 120, 14, 50, 5}
		min = numArr[0]
		max = numArr[0]
	)

	for _, value := range numArr {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}
	fmt.Println((max + min) + (max - min))
}
