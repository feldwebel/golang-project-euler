// euler0007 project main.go
/*
   What is the 10 001st prime number?

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
	"golang-project-euler/common"
)

func main() {
	n := 10001
	c := 0
	for i := 3; ; i += 2 {
		if i > 10 && i%10 == 5 {
			continue
		}
		if common.IsPrime(uint64(i)) {
			c++
			if c == n-1 {
				fmt.Println(i)
				break
			}
		}
	}
}
