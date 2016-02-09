// main
package main

import (
	"fmt"
	"math"
)

var result, i uint64

func speedupHack(n uint64) uint64 {
	return uint64(math.Sqrt(float64(n))) + 1
}

func isPrime(factor uint64) bool {
	lim := speedupHack(factor)
	var f uint64
	for f = 3; f < lim; f += 2 {
		if factor%f == 0 {
			return false
		}
	}
	return true
}

func main() {
	result = 2
	for i = 3; i < 2000000; i = i + 2 {
		if isPrime(i) {
			fmt.Println(i)
			result += i
		}
	}
	fmt.Println(result)
}
