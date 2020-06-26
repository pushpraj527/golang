package main

import (
	"fmt"
)

const a = 10

//another way to do this (untyped constant)

const (
	b = 12
	c = 43.2
	d = "raj"
)

// typed constant
const (
	e = 12
	f = 43.2
	g = "raj"
)

func main() {

	fmt.Println(a)
	fmt.Println("untyped constant")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("typed constant")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
