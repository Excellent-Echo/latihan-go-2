package main

import "fmt"

func deleteChars(str string) string {
	char := ""

	if len(str) <= 3 {
		fmt.Println("kata harus lebih dari 3 huruf")
	} else {
		for i := 1; i < len(str)-1; i++ {
			char += string(str[i])
		}
	}
	return char
}

func main() {
	fmt.Println(deleteChars("potong bebek angsa"))
}
