// euler0006 project main.go
/*
   Find the difference between the sum of the squares of the first one
   hundred natural numbers and the square of the sum.

   No bruteforce only Euler & Faulhaber formulae

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
)

func main() {
	n := 100
	square_of_sum := (1 + n) * n / 2             //Euler (1 + 2 + ... + n)
	sum_of_square := n * (n + 1) * (2*n + 1) / 6 //Faulhaber for 1^2 + ... +n^2
	fmt.Println(square_of_sum*square_of_sum - sum_of_square)
}
