package main

import "fmt"

func main() {
    name := "ronaldo"
    slice:= []
    
	
	for i := len(name) - 1; i >= 0; i-- {
        slice = append(slice, string(name[i]))
    }
    fmt.Printf(slice)
}