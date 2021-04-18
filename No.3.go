package main

import "fmt"

func main() {
	var angka = [...]int{20, 40, 50, 80}
	var max = angka[0]
	var min = angka[0]
	for _, angka := range angka {
		if angka > max {
			max = angka
		}
		if angka < min {
			min = angka
		}
	}
	fmt.Println((max + min) + (max - min))
}
