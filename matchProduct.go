package main

import "fmt"

func main() {
	var products = [...]int{10, 10, 20, 20, 10, 30, 50, 10, 20}
	var counter = 0

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
		for j := 0; i < len(products); i++ {
			if products[j] == products[i] {
				fmt.Println(products[j])
				counter++
			}
		}
	}

	fmt.Println(counter)
}
