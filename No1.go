package main

import "fmt"

func main() {
	var number = [5]int{1, 2, 3, 4, 5}
	var result int
	for i := 0; i < len(number); i++ {
		result += number[i]
		// fmt.Println(number[i])
	}
	fmt.Println(result)

}
