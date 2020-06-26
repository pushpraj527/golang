package main

import (
    "fmt"
)

var y = 21
var z int

func main() {
	
	//short declaration
	x := 42 // declaration and assignment
	fmt.Println(x)
	x = 72 // assignment here we use only "="
	fmt.Println(x)

	z = 27

	foo()
}

func foo() {
	fmt.Println("var variable is printed from foo:", y)
	fmt.Println("value of z:",z)
}