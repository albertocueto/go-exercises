// Creating a race condition by reading a variable and then doing soemthing with it while runtime.Gosched()

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

	for i := 0; i < numOfIncrements; i++ {
		go func() {
			incVal := value
			fmt.Println(value)
			runtime.Gosched()
			incVal++
			value = incVal
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(value)

}
