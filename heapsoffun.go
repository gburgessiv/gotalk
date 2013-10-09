package main

import (
	"fmt"
)

var nextInt = func() func() *int {
	i := 0
	return func() *int {
		i++
		next := i
		return &next
	}
}()

func main() {
	ints := []*int{}
	for i := 0; i < 100; i++ {
		ints = append(ints, nextInt())
	}

	for _, n := range ints {
		fmt.Println(*n)
	}
}
