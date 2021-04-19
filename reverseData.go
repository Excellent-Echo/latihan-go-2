package main

import "fmt"

func main() {
	var str = "impact byte"
	var sliceStr []string

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == " " {
			continue
		} else {
			sliceStr = append(sliceStr, string(str[i]))
		}
	}

	fmt.Println(sliceStr)
}
