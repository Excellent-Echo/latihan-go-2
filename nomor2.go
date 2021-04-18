package main

import "fmt"

func main() {
	var input string

	fmt.Print("masukkan input string untuk di reverse : ")
	fmt.Scanf("%s", &input)

	if len(input) >= 5 {
		var reverse string
		var result string

		for i := len(input) - 1; i >= 0; i-- {
			reverse += string(input[i])
		}
		result = reverse
		fmt.Println(result)
	} else {
		fmt.Println("Error")
	}
}
