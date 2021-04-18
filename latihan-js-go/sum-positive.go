package main

import "fmt"

func sumPositive(array []int) int {
	sum := 0

	for i := 0; i < len(array); i++ {
		if array[i] > 0 {
			sum += array[i]
		}
	}
	return sum
}

func main() {
	fmt.Println(sumPositive([]int{-1, -87, -25, -666})) // 0
	fmt.Println(sumPositive([]int{-4, 7, 5, -6}))       // 7 + 5 = 12
	fmt.Println(sumPositive([]int{1, -4, 7, 12}))       // 1 + 7 + 12 = 20
}
