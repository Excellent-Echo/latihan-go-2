package main

import "fmt"

func main() {
	data := [10]string{"a", "i", "u", "e", "e", "o", "c", "d", "f", "g"}
	result := make([]string, 0)
	middle := len(data) / 2

	if len(data)%2 == 0 {
		result = append(result, data[middle-1], data[middle])
	} else {
		result = append(result, data[middle-1], data[middle], data[middle+1])
	}
	fmt.Println(result)
}
