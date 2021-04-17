package main

import "fmt"

func main() {
	array := [...]int{10, 10, 20, 20, 10, 30, 50, 10, 20}
	// slice := []
	var tempArray []int
	var slice []int
	var slicex []int
	// var finalArray []int
	var otherArray []int

	for index, value := range array {
		if value == array[index+2] {
			tempArray = append(tempArray, value)
			slice = array[array[index]-len(array):]
		}

		// else {
		// 	otherArray = append(otherArray, value)
		// }

		if len(tempArray) > 2 {
			slicex = tempArray[:2]
		}

	}
	fmt.Println(slice)
	fmt.Println(slicex)
	fmt.Println(tempArray)
	fmt.Println(otherArray)
}
