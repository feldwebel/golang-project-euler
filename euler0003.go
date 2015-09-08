// euler project euler0003.go
package main

import (
    "fmt"
    "math"
    "time"
)
/*
It is mathematical fact: the prime divider cannot be greater than
dividend's square root
*/
func speedupHack(n int)int {
    return int(math.Sqrt(float64(n)))
}

func isPrime(factor int) bool {
    lim := speedupHack(factor)//near twofold speedup due to lim precomputation
    for f := 3; f < lim; f +=2 {
//        if f > 10 && f%10 == 5 {
//            continue //5% slowdown
//        }
        if factor % f == 0 {
            return false
        }
    }
    return true
}

func main(){
    start := time.Now()
    n := 600851475143
    number := speedupHack(n) //near no speedup

    for i:= 3; i <= number; i += 2 {
        if i > 10 && i%10 == 5 {
            continue //10% speedup
        }
        if isPrime(i) && n%i == 0 {
            fmt.Println(i) //Yes it is!
        }
    }
    
    fmt.Println("Time spent %s", time.Since(start))
}