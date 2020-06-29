package main

import (
	"fmt"
)

func main() {

	x := "praj"

	switch x {
	case "praj":
		fmt.Println("printed")
		fallthrough
	case "pushp":
		fmt.Println("Its not true but will be printed")

	default:
		fmt.Println("default is printed")
	}

	switch {
	case true:
		fmt.Println("printed")
		fallthrough
	case false:
		fmt.Println("Its not true but will be printed")
	case (1 == 1):
		fmt.Println("true but not be printed")
	default:
		fmt.Println("default is printed")
	}
}
