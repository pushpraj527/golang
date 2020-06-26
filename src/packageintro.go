package main

import (
    "fmt"
)

func main() {
    
	n, err := fmt.Println("Hello Raj", 27, true) // n and err is return by Println
	fmt.Println(n,err)

	n, _ = fmt.Println("Hello Raj", 27, true) //here "_" is use to take void input
	fmt.Println(n)

}