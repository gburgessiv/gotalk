package main

import "fmt"

func eatAll(q <-chan int) {
    ok := true
    for ok {
        _, ok = <-q
    }
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go eatAll(ch1)
    go eatAll(ch2)

    nch1 := 0
    nch2 := 0
    for i := 0; i < 10000; i++ {
        select {
        case ch1 <- 0:
            nch1++
        case ch2 <- 0:
            nch2++
        }
    }

    fmt.Println(nch1, nch2)
}
