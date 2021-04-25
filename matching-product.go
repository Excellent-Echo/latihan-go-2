package main

import "fmt"

func MatchingProduct(product []int)  {
	length := len(product)

	var one []int
	var two []int
	var three []int

	for i := 0 ; i < len(product) ; i++ {
		if product[0] == product[i] {
			//fmt.Println(product[i])
			one = append(one, product[i])
		} else if product[length - 1] == product[i]{
			//fmt.Println(product[i])
			two = append(two, product[i])
		} else if product[1] == product[i] {
			//fmt.Println(product[i])
			three = append(three, product[i])
		}
	}

	result := (len(one) / 2) + (len(two) / 2) + (len(three) / 2)
	fmt.Println(result)
}

func main() {
	product1 := []int{10, 10, 20, 20, 10, 30, 50, 10, 20}
	product2 := []int{1, 3, 3, 4, 5, 6, 1}
	product3 := []int{67, 213, 1, 0, 98, 2}
	MatchingProduct(product1)
	MatchingProduct(product2)
	MatchingProduct(product3)

}
