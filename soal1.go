package main

import "fmt"

func main() {
	arr := []int{2, 3, 6}

	var result int

	for _, num := range arr {
		result += num
	}
	fmt.Println(result)
}
