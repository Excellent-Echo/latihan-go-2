package main

import "fmt"

func main() {
	str := "impact byte"
	var array []string

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == string(" ") {
			continue
		} else {
			array = append(array, string(str[i]))
		}
	}
	fmt.Println(array)
}
