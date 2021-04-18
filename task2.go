package main

import "fmt"

func main() {
	var sample1 string = "radika"
	var result1 = make([]string, 0)

	for i := len(sample1) - 1; i >= 0; i-- {
		//fmt.Println(string(sample1[i]))
		//var result = [...]string{string(sample1[i])}
		result1 = append(result1, string(sample1[i]))
	}
	fmt.Println(result1)

	var sample2 string = "Radika Yudhistira Early"
	var result2 = make([]string, 0)

	for i := len(sample2) - 1; i >= 0; i-- {
		if string(sample2[i]) == string(" ") {
			continue
		} else {
			result2 = append(result2, string(sample2[i]))
		}
	}
	fmt.Println(result2)
}
