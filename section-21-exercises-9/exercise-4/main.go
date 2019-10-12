// Previous exercive reminder: Creating a race condition by reading a variable and then doing soemthing with it while runtime.Gosched()
// Now, we'll solve the race condition with mutex

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	value := 0
	var wg sync.WaitGroup
	numOfIncrements := 500
	wg.Add(numOfIncrements)

	var mutex sync.Mutex

	for i := 0; i < numOfIncrements; i++ {
		go func() {
			mutex.Lock()
			incVal := value
			fmt.Println(value)
			incVal++
			value = incVal
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(value)

}
