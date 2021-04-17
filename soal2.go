package main

import "fmt"

func main() {
	nama := "Randhika Rizkyaldi"
	slice := make([]string, 0)

	for i := len(nama) - 1; i >= 0; i-- {

		if string(nama[i]) == string(" ") {
			continue
		} else {
			slice = append(slice, string(nama[i]))
		}

	}

	fmt.Println(slice)
}
