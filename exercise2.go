package main

import "fmt"

func main() {
	var str string

	fmt.Printf("Masukan: ")
	fmt.Scanf("%s", &str)

	char := make([]string, 0)

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == (" ") {
			continue
		} else {
			char = append(char, string(str[i]), ",")
		}
	}
	fmt.Println(char)
}