package main

import "fmt"

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string

	fmt.Print("input word: ")
	fmt.Scanf("%s", &str)

	fmt.Println(isPalindrome(str))
}
