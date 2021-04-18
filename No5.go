package main

import "fmt"

func main() {
	var nama = [...]string{"j", "o", "m", "b", "l", "o", "i"}
	var hasil = make([]string, 0)
	var tengah = len(nama) / 2
	if len(nama)%2 == 0 {
		hasil = append(hasil, nama[tengah-1], nama[tengah])

	} else {
		hasil = append(hasil, nama[tengah-1], nama[tengah], nama[tengah+1])
	}

	fmt.Println(hasil)
}
