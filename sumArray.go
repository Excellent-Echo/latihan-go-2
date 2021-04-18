package main

import "fmt"

func main() {
	arr := []int{2, 4, 5, 6}

	var results int

	for _, sum := range arr {
		results = results + sum
	}
	fmt.Print(results)

}
