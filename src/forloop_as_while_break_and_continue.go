package main

import (
	"fmt"
)

func main() {

	x := 20

	for {

		if x < 8 {
			fmt.Println("Break will executed now")
			break
		}

		if x == 15 {
			fmt.Println("value of x and then continue", x)
			x--
			continue
		}

		fmt.Println("value of x:", x)
		x--
	}

}
