package main

import "fmt"

func fibs(c chan int64) {
	var a, b int64 = 1, 0

	for i := 0; i < 10; i++ {
		c <- a
		a, b = a+b, a
	}

	close(c)
}

func main() {
	ch := make(chan int64)
	go fibs(ch)
	for i := 0; i < 50; i++ {
		num, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Got fib", num) // ok glass
	}
}
