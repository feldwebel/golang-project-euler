// euler0001.go
/*
   Find the sum of all the multiples of 3 or 5 below 1000.

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
)

// Nothins serious, just straightforward division modulo N

func main() {

	var result int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	fmt.Println("The answer is", result)
}
