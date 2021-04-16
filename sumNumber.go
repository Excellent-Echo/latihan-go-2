package main

import "fmt"

func main() {
	// var number = [...]int{1, 2, 3, 4, 5}
	// var number = [...]int{1, 2, 3}
	var number = [...]int{}

	var result int
	for _, num := range number {
		result += num
	}
	fmt.Println(result)
}
