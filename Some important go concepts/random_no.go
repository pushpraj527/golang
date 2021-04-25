/* Generate random number based on seed value as time so that every
program starts we get different number*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(generateRandomNumber(1000))
	}

}

func generateRandomNumber(number int) (int, int) {
	s := rand.NewSource(time.Now().UnixNano()) //generating seed value based on time
	r := rand.New(s)                           //Here we are providing different seed value.
	x := r.Intn(number)                        // It will generate every time new number when we start a program

	y := rand.Intn(number) // It will generate first time same number when we start the program later on different

	return x, y
}

// Try to run this program 3-4 time you will get the difference between both
