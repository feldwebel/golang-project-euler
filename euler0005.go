// euler0005 project main.go
/*
   What is the smallest positive number that is evenly divisible
   by all of the numbers from 1 to 20?

   Prime factorization instead of straightforward bruteforce

   @author Dan Tereschenko <dan@tereschenko.info>
   @copyright 2015
*/
package main

import (
	"fmt"
	"math"
)

//func isFound(i int) bool {
//    for x := 11; x <= 20; x++ {
//        if i%x != 0 {
//            return false
//        }
//    }
//    return true
//}

//func main() {
//    for i := 20; ; i +=2 {
//        if isFound(i) {
//            fmt.Println(i)
//            break
//        }
//    }
//}

var prime = []int{2, 3, 5, 7, 11, 13, 17, 19}

func getFactors(n int) map[int]int {
	result := make(map[int]int)
	n1 := n
	for _, f := range prime {
		for {
			if remainder := n1 % f; remainder == 0 {
				n1 = n1 / f
				result[f]++
			} else {
				break
			}
		}
	}
	return result
}

func main() {
	r := 1
	factors := map[int]int{}
	for i := 11; i <= 20; i++ {
		v := getFactors(i)
		for key, val := range v {
			if factors[key] < val {
				factors[key] = val
			}
		}
	}
	for key, val := range factors {
		r = r * int(math.Pow(float64(key), float64(val)))
	}
	fmt.Println(r)
}
