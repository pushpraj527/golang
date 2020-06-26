package main

import (
	"fmt"
)

var a int

type raj int

var b raj

func main() {

	a = 21
	b = 527

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b) //conversion, not type casting

	fmt.Println("value of a: ",a)

}
