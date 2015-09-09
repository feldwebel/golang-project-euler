// euler0005 project main.go
/*
   What is the smallest positive number that is evenly divisible 
   by all of the numbers from 1 to 20?

   Near straightforward brutforce

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
)

func isFound(i int) bool {
	for x := 11; x <= 20; x ++ {// skip 1-10 as dividers in 11-20
		if i%x != 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 20; ; i += 2 {
		if isFound(i) {
			fmt.Println(i)
			break
		}
	}
}
