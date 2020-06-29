package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {

			a := 4
			for a > 0 {
				fmt.Println("value of a: ", a)
				a--
			}
			fmt.Println("value of j:", j)
		}
		fmt.Println("value of i:", i)
	}

}
