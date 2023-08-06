package main

import (
	"fmt"
	"sync"
)

type SimpleRWMutex struct {
	mu      sync.Mutex
	readers int
	writers int
}

// command + L to select lines
func (rw *SimpleRWMutex) RLock() {
	rw.mu.Lock()
	rw.readers++
	rw.mu.Unlock()
}

func (rw *SimpleRWMutex) RUnlock() {
	rw.mu.Lock()
	rw.readers--
	rw.mu.Unlock()
}

func (rw *SimpleRWMutex) Lock() {
	rw.mu.Lock()
	for rw.readers > 0 || rw.writers > 0 {
		rw.mu.Unlock()
		rw.mu.Lock()
	}
	rw.writers++
	rw.mu.Unlock()
}

func (rw *SimpleRWMutex) Unlock() {
	rw.mu.Lock()
	rw.writers--
	rw.mu.Unlock()
}

func main() {
	// RWMutex
	fmt.Println("Running 100 concurrent go routines updating counter with custom mutex implemantation")

	var (
		rwMutex SimpleRWMutex
		wg      sync.WaitGroup
	)

	counter := 0

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			rwMutex.RLock()
			fmt.Println("Current value: ", counter)
			rwMutex.RUnlock()

			rwMutex.Lock()
			counter++
			rwMutex.Unlock()
		}()

		wg.Wait()

		rwMutex.Lock()
		fmt.Println("Final value: ", counter)
		rwMutex.Unlock()
	}
}
