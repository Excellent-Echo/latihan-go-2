package main

import "fmt"

func main() {
	name := "Teuku Muhammad Ammar"
	slice := make([]string, 0)
	length := len(name)

	for i := length - 1; i >= 0; i-- {
		if string(name[i]) != string(" ") {
			slice = append(slice, string(name[i]))
		} else {
			continue
		}
	}
	fmt.Print(slice)
}
