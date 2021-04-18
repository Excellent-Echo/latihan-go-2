package main

import "fmt"

func main() {
	data := [10]string{"H", "e", "l", "m", "i", "y", "u", "s", "u", "f"}
	result := make([]string, 0)
	middle := len(data) / 2

	if len(data)%2 == 0 {
		result = append(result, data[middle-1], data[middle])
	} else {
		result = append(result, data[middle-1], data[middle], data[middle+1])
	}
	fmt.Println(result)
}
