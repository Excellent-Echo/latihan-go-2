package main

import (
	"fmt"
	"math"
)

func changeStringPosition(str string) string {
	lastString := ""
	firstString := ""
	midString := ""
	evenString := float64(len(str)) / 2
	oddString := len(str)%2 != 0

	if oddString {
		midString = string(str[int(evenString)])
	}

	for i := 0; i <= int(evenString)-1; i++ {
		firstString += string(str[i])
	}

	for i := int(math.Ceil(float64(evenString))); i < len(str); i++ {
		if oddString {
			lastString += string(str[i])
		} else {
			lastString += string(str[i])
		}
	}
	return lastString + midString + firstString

}

func main() {
	fmt.Println(changeStringPosition("aaabccc"))        // cccbaaa
	fmt.Println(changeStringPosition("aab"))            // baa
	fmt.Println(changeStringPosition("aaaacccc"))       // ccccaaaa
	fmt.Println(changeStringPosition("abcdefghabcdef")) // habcdefabcdefg
}
