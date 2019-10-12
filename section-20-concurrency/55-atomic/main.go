package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goRoutines:", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)                         // Write to the counter safely
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) // Read the counter safely
			// time.Sleep(time.Second)
			runtime.Gosched() // Hey, go do something else. Cleaner.
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait() // Don't exit my program until everything's done.
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
