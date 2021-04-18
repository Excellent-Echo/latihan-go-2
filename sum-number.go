package main

import "fmt"

func main() {
	
	var (
		num =[]int{1, 2, 3 ,4 ,5}
	 	total int
	)
	
	for _,num:= range num{
	
		total +=num
	}
	fmt.Println(total)
///////////////////////////////////////
	var (
		num1 =[]int{1, 2, 3}
		total1 int
	)
	
	for _,num1:= range num1{
	
		total1 +=num1
	}
	fmt.Println(total1)
/////////////////////////////////////
	var (
		num2 =[]int{}
		total2 int
	)
	
	for _,num2:= range num2{
		total2 +=num2
	}
	
	fmt.Println(total2)
}




