package main

import (
	"fmt"
	"runtime"
	"time"
)

func quickSort(n []int, wait chan bool) {
	const SPLIT = 25000
	if len(n) < 2 {
		return
	}

	midAt := len(n) / 2
	pivot := n[midAt]
	n[0], n[midAt] = n[midAt], n[0]
	mid := 1
	for i, q := range n[1:] {
		if q <= pivot {
			n[mid], n[i+1] = n[i+1], n[mid]
			mid++
		}
	}

	var wait1, wait2 chan bool
	if mid > SPLIT {
		wait1 = make(chan bool)
		go quickSort(n[:mid], wait1)
	} else {
		quickSort(n[:mid], nil)
	}

	if len(n)-mid > SPLIT {
		wait2 = make(chan bool)
		go quickSort(n[mid:], wait2)
	} else {
		quickSort(n[mid:], nil)
	}

	if wait1 != nil {
		<-wait1
	}
	if wait2 != nil {
		<-wait2
	}
	if wait != nil {
		close(wait)
	}
}

func isSorted(arr []int) bool {
	l := arr[0]
	for _, n := range arr[1:] {
		if n < l {
			return false
		}
		l = n
	}
	return true
}

func main() {
	const NUM = 100000000
	const PROCS = 4
	nums := make([]int, NUM)
	for i := range nums {
		nums[i] = NUM - i
	}

	cb := make(chan bool)
	runtime.GOMAXPROCS(PROCS)
	start := time.Now()
	quickSort(nums, cb)
	<-cb
	diff := time.Since(start)

	if isSorted(nums) {
		fmt.Println("Sorted! Took", diff)
	} else {
		fmt.Println("Not sorted")
	}
}
