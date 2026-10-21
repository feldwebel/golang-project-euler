// euler0016.go
/*
   Find the sum of all the digits of 2^1000.

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
	x := new(big.Int).Lsh(big.NewInt(1), 1000) // 1 << 1000
	y := x.String()

	for i := 0; i < len(y); i++ {
		b := y[i]        // byte, значение '0'..'9'
		digit := b - '0' // 0..9
		sum += int(digit)
	}
	fmt.Println(sum)
}
