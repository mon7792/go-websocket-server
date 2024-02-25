package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Start a few Goroutines for demonstration

	for i := 0; i < 5; i++ {
		go func(num int) {
			startTime := time.Now()
			defer func() {
				fmt.Printf("GR: %d Time: %d", num, time.Since(startTime))
				fmt.Println()
			}()
			time.Sleep(time.Duration(num) * time.Millisecond)
		}(i)
	}

	time.Sleep(5 * time.Second)

	// Get the number of running Goroutines
	numGoroutines := runtime.NumGoroutine()

	fmt.Printf("Number of Running Goroutines: %d\n", numGoroutines)
}
