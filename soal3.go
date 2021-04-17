package main

import "fmt"

func main() {
	arrya := [5]int{10, 120, 14, 50, 5}
	var max, min, sum, sub int

	max = arrya[0]
	min = arrya[0]

	for i := 0; i < len(arrya); i++ {

		if arrya[i] > max {
			max = arrya[i]
		} else if arrya[i] < min {
			min = arrya[i]
		}
	}

	sum = max + min
	sub = max - min

	fmt.Println(max, " + ", min, " = ", sum)
	fmt.Println(max, " - ", min, " = ", sub)
}
