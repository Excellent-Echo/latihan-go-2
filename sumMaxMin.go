package main

import (
	"fmt"
	"math"
)

func main() {
	// var number = [...]int{10, 120, 14, 50, 5}
	var number = [...]int{0, 2, 3, 4, 5}
	var max = 0
	var min = math.MaxInt64

	for _, num := range number {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	var MaxPlusMin = max + min
	var MaxMinusMin = max - min
	var result = MaxPlusMin + MaxMinusMin

	fmt.Println(result)
}
