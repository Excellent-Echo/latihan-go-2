package main

import "fmt"

func main() {
	var sample1 = [5]int8{1, 2, 3, 4, 5}
	//var sample2 = [3]int8{1, 2, 3}
	var counter int8

	for i := 0; i <= len(sample1)-1; i++ {
		//fmt.Println(sample1[i])
		counter += sample1[i]
	}
	fmt.Println(counter)

	// for i := 0; i <= len(sample2)-1; i++ {
	// 	counter += sample2[i]
	// }
	// fmt.Println(counter)
}
