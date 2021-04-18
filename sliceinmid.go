package main

import "fmt"

func main(){
	var(
		data      = [...]string{"a", "n", "a", "n", "s", "y", "a"} 
		midValue  = len(data) / 2
		hasil	  = make([]string,0)
	)

	

	if len(data)%2 == 0 {
		hasil = append(hasil, data[midValue-1], data[midValue])
	} else {
		hasil = append(hasil, data[midValue-1], data[midValue], data[midValue+1])
	}
	fmt.Println(hasil)

}