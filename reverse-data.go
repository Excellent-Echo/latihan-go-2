package main

import "fmt"

func main() {
    var (
        name string = "ronaldo"
        len         = len(name)
        reverse []string
    ) 
    
    
    
	for i := len - 1; i >= 0; i-- {
        if string(name[i]) == string(""){

        }else{
            reverse=append(reverse, string(name[i]))
            continue

        }
    }
    
    fmt.Printf("%s", reverse)
}