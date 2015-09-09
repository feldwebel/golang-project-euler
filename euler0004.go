// euler0004
/*
   Find the largest palindrome made from the product of two 3-digit numbers.

   Keep in mind the multiplication is commutative, so there is no need check
   998x999 having 999x998 checked

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool { //some pain due to typecasting
	s := strconv.Itoa(n)
	l := len(s)

	for i := 0; i < int(l/2); i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var r, result int

	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- { //multiplication is commutative
			r = i * j
			if isPalindrome(r) && r > result {
				result = r
			}
		}
	}
	fmt.Println(result)
}