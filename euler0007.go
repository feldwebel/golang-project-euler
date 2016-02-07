// euler0007 project main.go
/*
   What is the 10 001st prime number?

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
	"math"
)

func speedupHack(n int) int {
	return int(math.Sqrt(float64(n))) + 1
}

func isPrime(factor int) bool {
	lim := speedupHack(factor)
	for f := 3; f < lim; f += 2 {
		if factor%f == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 10001
	c := 0
	for i := 3; ; i += 2 {
		if i > 10 && i%10 == 5 {
			continue
		}
		if isPrime(i) {
			c++
			if c == n-1 {
				fmt.Println(i)
				break
			}
		}
	}
}
