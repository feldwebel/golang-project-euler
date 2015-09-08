// euler0001.go
package main

import (
	"fmt"
)

func main() {

	var result int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	fmt.Println("The answer is", result)
}
