package main

import "fmt"

func main() {
	// array

	// deklarasi awal
	// var names [4]string

	// names[0] = "afis"

	// // langsung diisi dengan data
	// var nameMentor = [3]string{
	// 	"andi",
	// 	"afis",
	// 	"eddy",
	// }

	// [...] three dots, literal
	// var number = [...]int{
	// 	1,
	// 	2,
	// 	3,
	// 	4,
	// 	5,
	// 	6,
	// 	7,
	// }

	// number[6] = 8

	// for range untuk array

	// names := [3]string{
	// 	{{"afis"}, "pratama"},
	// 	{"mas", "eddy"},
	// 	{"mas", "david"},
	// }

	// for _, name := range names {

	// }

	var number = [5]int{1, 2, 3, 4, 5}

	slice1 := number[:3]
	slice2 := number[3:]
	slice3 := number[1:4]
	slice4 := number[:]

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))

	// append slice

	// var number = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// var sliceNum = number[1:4]

	// slice := append(sliceNum, 11, 12, 13)

	// // var slice []int

	// fmt.Println(slice)

	// make slice

	// var slice = make([]int, 5, 10)

	// slice[0] = 20

	// slice2 := append(slice, 10, 11, 12)

	// fmt.Printf("data %v dengan tipe %T", slice2, slice2)

	// var number = [5]int{1, 2, 3, 4, 5}

	// var fromSlice = number[:]
	// var toSlice = make([]int, len(fromSlice), cap(fromSlice))

	// copy(toSlice, fromSlice)

	// fmt.Println(toSlice, fromSlice)

}
