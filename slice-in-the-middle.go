package main

import (
	"fmt"
	"strings"
)

func main() {
	data := [5]string{"a", "i", "u", "e", "o"}

	satu := data[len(data)/2|0]
	dua := data[len(data)/3|0]
	tiga := data[len(data)/2|1]

	values := []string{}

	if len(data)%2 == 0 {
		values = append(values, satu)
		values = append(values, dua)
		result := strings.Join(values, ",")
		toArray := strings.Fields(result)
		fmt.Println(toArray)
	} else {
		values = append(values, dua)
		values = append(values, satu)
		values = append(values, tiga)
		result := strings.Join(values, ",")
		toArray := strings.Fields(result)
		fmt.Println(toArray)
	}
}
