// main
package main

import (
	"fmt"

	"github.com/feldwebel/golang-project-euler/common"
)

var result, i uint64

func main() {
	result = 2
	for i = 3; i < 2000000; i = i + 2 {
		if common.IsPrime(i) {
			result += i
		}
	}
	fmt.Println(result)
}
