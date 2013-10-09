package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func bitonicComp(a, b *int) {
	if *a >= *b {
		*b, *a = *a, *b
	}
}

func findPowerOfTwo(n int) int {
	var i int
	for i = 0; n != 0; i++ {
		n >>= 1
	}
	return i
}

func bitonicSort(array []int) {
	lock := &sync.Mutex{}
	cond := sync.NewCond(lock)
	numToWait := 0

	incSig := func(howMany int) {
		lock.Lock()
		numToWait += howMany
		lock.Unlock()
	}

	decSig := func(howMany int) {
		incSig(-howMany)
		if numToWait <= 0 {
			lock.Lock()
			numToWait = 0
			cond.Broadcast()
			lock.Unlock()
		}
	}

	blockOrSig := func() {
		lock.Lock()
		numToWait -= 1
		if numToWait == 0 {
			cond.Broadcast()
		} else {
			cond.Wait()
		}
		lock.Unlock()
	}

	var bitonicSliceSecond func([]int, bool)
	bitonicSliceSecond = func(array []int, wait bool) {
		jmp := len(array) / 2
		start := 0
		end := start + jmp
		for end < len(array) {
			bitonicComp(&array[start], &array[end])
			start++
			end++
		}

		if len(array) > 10000 {
			incBy := 1
			if wait {
				incBy++
			}
			incSig(1)
			go bitonicSliceSecond(array[:len(array)/2], true)
			go bitonicSliceSecond(array[len(array)/2:], true)
		} else {
			if len(array) > 2 {
				bitonicSliceSecond(array[:len(array)/2], false)
				bitonicSliceSecond(array[len(array)/2:], false)
			}

			if wait {
				decSig(1)
			}
		}
	}

	bitonicSliceFirst := func(array []int) {
		start := 0
		end := len(array) - 1
		for end > start {
			bitonicComp(&array[start], &array[end])
			start++
			end--
		}

		if len(array) > 10000 {
			incSig(1)
			go bitonicSliceSecond(array[:len(array)/2], true)
			go bitonicSliceSecond(array[len(array)/2:], true)
		} else {
			if len(array) > 2 {
				bitonicSliceSecond(array[:len(array)/2], false)
				bitonicSliceSecond(array[len(array)/2:], false)
			}
			decSig(1)
		}
	}

	for i := 2; i <= len(array); i <<= 1 {
		incSig(len(array)/i + 1)
		for s := 0; s < len(array); s += i {
			go bitonicSliceFirst(array[s : s+i])
		}
		blockOrSig()
	}
}

func isSorted(array []int) bool {
	if len(array) < 2 {
		return true
	}

	old := array[0]
	for _, n := range array[1:] {
		if n < old {
			return false
		}
		old = n
	}
	return true
}

func parallelIsSorted(array []int, numProcs int) bool {
	step := len(array) / numProcs
	chans := []chan bool{}
	for i := 0; i < numProcs; i++ {
		c := make(chan bool, 1)
		chans = append(chans, c)
		go func(n int) {
			start := n * step
			end := (n + 1) * step
			if n > 0 {
				start--
			}
			if n+1 < numProcs {
				end++
			}
			c <- isSorted(array[start:end])
			close(c)
		}(i)
	}

	for _, c := range chans {
		if !<-c {
			return false
		}
	}
	return true
}

func main() {
	const PROCS = 4
	const NUM_ELEMENTS = 1 << 20
	runtime.GOMAXPROCS(4)
	array := make([]int, NUM_ELEMENTS)
	for i := 0; i < NUM_ELEMENTS; i++ {
		array[i] = NUM_ELEMENTS - i
	}

	start := time.Now()
	bitonicSort(array)
	diff := time.Since(start)

	if parallelIsSorted(array, PROCS) {
		fmt.Println("Array is sorted in", diff)
	} else {
		fmt.Println("NOOOOOOOOOOOOOOOO")
		for _, n := range array {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}
}
