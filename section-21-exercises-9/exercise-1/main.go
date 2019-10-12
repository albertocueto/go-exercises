package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go process1(&wg)
	go process2(&wg)
	wg.Wait()
}

func process1(wg *sync.WaitGroup) {
	fmt.Println("Ground control to major Tom.")
	wg.Done()
}

func process2(wg *sync.WaitGroup) {
	fmt.Println("This is major Tom to ground control.")
	wg.Done()
}
