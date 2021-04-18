package main

import "fmt"

func main() {
	str := "Danang"
	var newArray []string

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == string(" ") {
			continue
		} else {
			newArray = append(newArray, string(str[i]))
		}
	}
	fmt.Println(newArray) // [g n a n a D]
}