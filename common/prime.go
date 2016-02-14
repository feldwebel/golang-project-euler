package common

import (
	"math"
)

func speedupHack(n uint64) uint64 {
	return uint64(math.Sqrt(float64(n))) + 1
}

func IsPrime(factor uint64) bool {
	lim := speedupHack(factor)
	var f uint64
	for f = 3; f <= lim && factor%f != 0; f += 2 {
	}

	return factor == 2 || f >= lim
}
