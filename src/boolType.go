package main

import (
	"fmt"
)

var a bool

func main() {

	fmt.Println("Zero value for 'a' :", a)

	x := 10
	y := 20

	fmt.Println("x: ", x, "y: ", y)

	fmt.Println("x == y : ", x == y)

	fmt.Println("x !=  y : ", x != y)

	fmt.Println("x < y :", x < y)

}
