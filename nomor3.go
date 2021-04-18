package main

import "fmt"

func main() {
	var (
		max, min, summax, summin int
	)

	array := [5]int{10, 120, 14, 50, 5}
	max = array[0]
	min = array[0]
	for i := 1; i <= len(array)-1; i++ {
		if array[i] > max {
			max = array[i]
		} else if array[i] < min {
			min = array[i]
		}
	}
	summax = max + min
	summin = max - min
	fmt.Printf("hasil dari summax %v + %v = %v, dan hasil dari summin %v - %v = %v", max, min, summax, max, min, summin)
}
