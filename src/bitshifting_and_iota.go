package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	x := 4
	fmt.Printf(" %d\t\t\t%b\n", x, x)
	fmt.Println("after bit shifting")
	y := x << 1
	fmt.Printf("%d\t\t\t%b\n", y, y)

	fmt.Println("Bitshift and iota")
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
