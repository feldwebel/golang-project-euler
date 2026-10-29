/*
   Find the sum of all the digits of 100!.

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2026
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := 0
	result := new(big.Int).MulRange(1, 100)
	y := result.String()

	for i := 0; i < len(y); i++ {
		b := y[i]        // byte, значение '0'..'9'
		digit := b - '0' // 0..9
		sum += int(digit)
	}
	fmt.Println(sum)
	fmt.Println(result)
}
