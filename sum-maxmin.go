package main

import "fmt"

func main(){

var (
	angka = [...]int{16, 12, 7,8, 4}
	min= angka[0]
	max= angka[0]
)

for _,angka := range angka {
	if angka > max {
		max = angka
	}

	if angka < min{
		min = angka
	}
}

fmt.Println(angka,"(", max,"+",min,")", "+", "(", max,"-",min,")", "=", (max+min), "+", (max-min), "=",
(max+min)+(max-min))

}