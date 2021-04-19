package main

import "fmt"

func main() {
	str := "impact byte"
	char := make([]string, 0)

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == (" ") {
			continue
		} else {
			char = append(char, string(str[i]))
		}
	}
	fmt.Println(char)
}
