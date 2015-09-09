// euler0002 project main.go
/*
   By considering the terms in the Fibonacci sequence whose values
   do not exceed four million, find the sum of the even-valued terms.

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
)

//AFAIK there are some other problems with Fibonacci sequense wait me ahead.
//Hence I try to separate its calculation in struct to package it in the future

type Fib struct {
	prev, curr uint64 //Fibonacci numbers are positive only
}

func (f *Fib) next() { //pass by reference used, working with attributes directly
	f.prev, f.curr = f.curr, f.curr+f.prev
}

func (f Fib) EvenSum(limit uint64) uint64 {
	var result uint64
	for f.curr < limit {
		if f.curr%2 == 0 {
			result += f.curr
		}
		f.next()
	}
	return result
}

func main() {
	f := Fib{1, 1}              //unfamiliar Go-styled OOP
	fmt.Println(f.EvenSum(4e6)) //four million
}
