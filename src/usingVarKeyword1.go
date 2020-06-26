package main

import (
    "fmt"
)

var y = 21


func main() {
	
	//short declaration
	x := 42 // declaration and assignment
	fmt.Println(x)
	x = 72 // assignment here we use only "="
	fmt.Println(x)

	foo()
}

func foo() {
	fmt.Println("var variable is printed from foo:", y)
}