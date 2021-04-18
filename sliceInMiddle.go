package main

import "fmt"

func main() {
	array := []string{"a", "i", "u", "e", "o"}

	length := len(array)
	lengthDiv := length / 2

	genap := array[lengthDiv-1 : lengthDiv+1]
	ganjil := array[lengthDiv-1 : length-1]

	if length%2 != 0 {
		fmt.Print(ganjil)
	} else {
		fmt.Print(genap)
	}

}
