package main

import "fmt"

func main() {
	var sample1 = [5]int16{10, 120, 14, 50, 5}
	var max, min, sum, sub, res int16

	max = 0
	min = 1000

	for i := 0; i <= len(sample1)-1; i++ {
		if sample1[i] > max {
			max = sample1[i]
		} else if sample1[i] < min {
			min = sample1[i]
		}
	}
	sum = max + min
	sub = max - min
	res = sum + sub

	fmt.Printf("(%d + %d) + (%d - %d) = %d + %d = %d ", max, min, max, min, sum, sub, res)
}
