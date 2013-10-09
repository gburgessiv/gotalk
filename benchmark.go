package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const NUM = 50000
	ch := make(chan int)
	runtime.GOMAXPROCS(1) // Force 1 process, tops.
	fmt.Println("Brb, spawning", NUM, "goroutines")

	memStatsNow := runtime.MemStats{}
	memStatsAfter := runtime.MemStats{}
	runtime.ReadMemStats(&memStatsNow)
	startTime := time.Now()
	for i := 0; i < NUM; i++ {
		go func(num int) { ch <- num }(i)
	}
	duration := time.Since(startTime)

	fmt.Println("Done")
	fmt.Println("To prove it, we're running", runtime.NumGoroutine(), "goroutines")
	runtime.ReadMemStats(&memStatsAfter)
	diff := memStatsAfter.Alloc - memStatsNow.Alloc
	fmt.Println("That took", duration)
	fmt.Println("...Which means an average of", duration/(NUM), "per goroutine")
	fmt.Println("Also,", NUM, "goroutines takes", diff, "bytes")
	fmt.Println("...Which leaves approx", int(diff/NUM), "bytes allocated for each gorotuine. Tiny, eh?")

	fmt.Println("\nThat's cool. Let's see how long it takes to swap the currently running goroutine", NUM*2, "times.")
	fmt.Println("In the interest of transparency, GOMAXPROCS =", runtime.GOMAXPROCS(0))

	startTime2 := time.Now()
	for i := 0; i < NUM; i++ {
		<-ch
	}
	duration2 := time.Since(startTime2)

	fmt.Println("Took", duration2)
	fmt.Println("...Which means", duration2/(2*NUM), "per goroutine swap")
	fmt.Println("Pretty good, right?")
}
