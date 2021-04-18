package main

import (
	"fmt"
)

func reverse(str string) string {
	array := []string{}
	temp := ""
	result := ""

	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			array = append(array, temp)
			temp = ""
		} else {
			temp += string(str[i])
		}
	}

	if len(temp) >= 0 {
		array = append(array, temp)
	}

	for j := len(array) - 1; j >= 0; j-- {
		result += " " + array[j]
	}
	return result
}

func main() {
	// fmt.Print(reverse("Hello World"))
	fmt.Print(reverse("Lorem Ipsum Dolor Sit Amet"))
}
