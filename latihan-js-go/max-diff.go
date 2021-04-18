package main

import "fmt"

func maxDiff(array []int) int {
	largest_number := array[0]
	smallest_number := array[0]

	for i := 0; i < len(array); i++ {
		if array[i] > largest_number {
			largest_number = array[i]
		} else {
			smallest_number = array[i]
		}
	}
	return largest_number - smallest_number
}

func main() {
	fmt.Println(maxDiff([]int{2, 4, 6, 8, 10}))       // 10 - 2 = 8
	fmt.Println(maxDiff([]int{7, 2, 19, 46, -3}))     // 46 - (-3) = 49
	fmt.Println(maxDiff([]int{-2, 0, -33, 55, -200})) // 55 - (-200) = 255
	fmt.Println(maxDiff([]int{-1, -3, -5}))           // (-1) - (-5) = 4
}
