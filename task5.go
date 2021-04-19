package main

import "fmt"

func main() {
	sample1 := [5]string{"a", "i", "u", "e", "o"}
	res := make([]string, 0)
	mid := len(sample1) / 2

	if len(sample1)%2 == 0 {
		//genap
		res = append(res, sample1[mid-1], sample1[mid])
	} else {
		//ganjil
		res = append(res, sample1[mid-1], sample1[mid], sample1[mid+1])
	}
	fmt.Println(res)
}
