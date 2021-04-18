package main

import "fmt"

func reverseWord(nama string) {
	
	for i := len(nama) - 1; i >= 0; i-- {
		fmt.Printf("%c", nama[i])
	}

}

func main() {
	var inputtedName string
	fmt.Print("enter your name:")
	fmt.Scanf("%s", &inputtedName)
	reverseWord(inputtedName)
}
