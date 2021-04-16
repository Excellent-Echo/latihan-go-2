package main

import (
	"fmt"
)

func main() {
	// data := [...]string{"a", "i", "u", "e", "o"}
	data := [...]string{"j", "o", "m", "b", "l", "o"}

	x := len(data) / 2

	even := data[x-1 : x+1]
	odd := data[x-1 : len(data)-1]

	if len(data)%2 == 0 {
		fmt.Println(even)
	} else {
		fmt.Println(odd)
	}

	// for _, num := range data {
	// fmt.Println(slice)
	// }
}
