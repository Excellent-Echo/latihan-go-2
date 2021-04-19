package main

import "fmt"

func main() {
	var (
		data = [...]string{"a", "i", "u", "e", "o"}
		division = len(data) / 2
		even = data[division - 1 : division + 1]
		odd = data[division - 1 : len(data) - 1]
	)

	if len(data) % 2 == 0 {
		fmt.Println(even)
	} else {
		fmt.Println(odd)
	}
}
