// euler0009 project main.go
/*
    There exists exactly one Pythagorean triplet for which a + b + c = 1000.
    Find the product abc.

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2016
*/
package main

import (
	"fmt"
)

const max = 1000

var a, b, c uint

func main() {
	for c = 1; c <= max-2; c++ {
		for b = 1; b < (max-c-1)>>1; b++ {
			a = max - b - c
			if a*a+b*b == c*c {
				fmt.Println(a, b, c, a*b*c)
			}
		}
	}
}
