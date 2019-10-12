// Previous exercive reminder: Creating a race condition by reading a variable and then doing soemthing with it while runtime.Gosched()
// Now, we'll solve the race condition with package atomic

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var value int64
	value = 0
	var wg sync.WaitGroup
	numOfIncrements := 500
	wg.Add(numOfIncrements)

	for i := 0; i < numOfIncrements; i++ {
		go func() {
			atomic.AddInt64(&value, 1)
			fmt.Println(atomic.LoadInt64(&value))
			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(value)
}
