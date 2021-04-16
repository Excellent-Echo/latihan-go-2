package main

import (
	"fmt"
)

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	result := 0

	for _, num := range array {
		result += num
	}
	fmt.Println(result)
}
