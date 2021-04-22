package main

import "fmt"

func main() {
	data := [...]string{"a", "i", "u", "e", "o"}

	switch len(data) % 2 == 0 {
	case true:
		fmt.Println(data[(len(data) / 2) - 1 : (len(data) / 2) + 1]) //genap
	case false:
		fmt.Println(data[(len(data) / 2) - 1 : len(data) - 1]) //ganjil
	}
}
