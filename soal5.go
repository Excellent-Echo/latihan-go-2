package main

import "fmt"

func main() {
	var kata = [...]string{"a", "b", "c", "d", "e", "f"}
	var hasil = make([]string, 0)
	var bagi = len(kata) / 2

	if len(kata)%2 == 0 {
		hasil = append(hasil, kata[bagi-1], kata[bagi])
	} else {
		hasil = append(hasil, kata[bagi-1], kata[bagi], kata[bagi+1])
	}
	fmt.Print(hasil)
}
