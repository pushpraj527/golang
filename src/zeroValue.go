package main

import (
    "fmt"
)

var a int
var b string
var c float32

func main() {

	fmt.Println("Zero Value")
	fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    a = 2
    b = "raj"
    c = 3.2
	fmt.Println("After assignment with some value")
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

}