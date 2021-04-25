package main

import (
	"fmt"
	"strings"
)

func ReverseData(name string) {
	var newName string

	for i := len(name) - 1; i >= 0; i-- {
		newName += string(name[i])
	}
	toArray := strings.Fields(newName)
	fmt.Println(toArray)
}

func main() {
	ReverseData("afista")
	ReverseData("skilvul")
	ReverseData("impact byte")
}
