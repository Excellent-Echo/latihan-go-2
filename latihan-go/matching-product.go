package main

import "fmt"

func main() {
	// array := [...]int{10, 10, 20, 20, 10, 30, 50, 10, 20} // 5 ??
	array := [...]int{1, 3, 3, 4, 5, 6, 1} // 2
	// array := [...]int{67, 213, 1, 0, 98, 2} // 0

	count := 0

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				count++
				break
			}
		}
	}

	fmt.Println("output adalah", count)
}
