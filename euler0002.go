// euler0002 project main.go
package main

import (
	"fmt"
)

type Fib struct{
    prev, curr uint64
}

func (f *Fib) next() {
    f.prev, f.curr = f.curr, f.curr + f.prev
}

func (f Fib) EvenSum(limit uint64) uint64 {
    var result uint64;
    for f.curr < limit {
        if f.curr%2 == 0 {
            result += f.curr
        }
        f.next()
    }
    return result
}

func main() {
    f := Fib{1,1}
    fmt.Println(f.EvenSum(4e6))
}
