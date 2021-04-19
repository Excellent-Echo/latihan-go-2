package main

import "fmt"

func main() {
	// var data = [5]string{"a", "i", "u", "e", "o"}
	// var data = [4]string{"b", "a", "c", "a"}
	// var data = [6]string{"j", "o", "m", "b", "l", "o"}
	var data = [5]string{"k", "e", "r", "e", "n"}

	if len(data)%2 == 0 {
		result := data[(len(data)/2 - 1):(len(data)/2 + 1)]
		fmt.Println(result)
	} else {
		result := data[(len(data)/2 - 1):(len(data)/2 + 2)]
		fmt.Println(result)
	}
}
