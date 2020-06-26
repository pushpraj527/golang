package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	d = iota + 1
	e = iota
	f = iota
	g
	h = iota
)

func main() {

	fmt.Println("value of a:", a)
	fmt.Println("value of b:", b)
	fmt.Println("value of c:", c)
	fmt.Println("second constant values")
	fmt.Println("value of d:", d)
	fmt.Println("value of e:", e)
	fmt.Println("value of f:", f)
	fmt.Println("value of g:", g)
	fmt.Println("value of h:", h)
}
