package main

import "fmt"

func main() {
	arrya := []int{1, 2, 3, 4, 5}

	var jml int

	for _, val := range arrya {
		jml += val
	}

	fmt.Println(jml)
}
