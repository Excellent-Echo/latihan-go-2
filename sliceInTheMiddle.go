package main

import "fmt"

func main() {

	//even
	array1 := [...]string{"D", "a", "n", "a", "n", "g"}
	result1 := make([]string, 0, 0)
	middle1 := len(array1) / 2

	if len(array1)%2 == 0 {
		result1 = append(result1, array1[middle1-1], array1[middle1]) // [n, a] 
	} else {
		result1 = append(result1, array1[middle1-1], array1[middle1], array1[middle1+1])
	}
	fmt.Println(result1)

	//odd
	array2 := [...]string{"E", "s", "t", "u", "t", "o", "m", "o", "a", "j", "i"}
	result2 := make([]string, 0, 0)
	middle2 := len(array2) / 2

	if len(array2)%2 == 0 {
		result2 = append(result2, array2[middle2-1], array2[middle2]) // [t, o, m] 
	} else {
		result2 = append(result2, array2[middle2-1], array2[middle2], array2[middle2+1])
	}
	fmt.Println(result2)
}