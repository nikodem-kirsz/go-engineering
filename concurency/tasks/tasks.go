package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg      sync.WaitGroup
		counter int
		rwMutex sync.RWMutex
	)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			rwMutex.RLock()
			fmt.Printf("Incrementing current counter: %v\n", counter)
			rwMutex.RUnlock()

			rwMutex.Lock()
			counter++
			rwMutex.Unlock()

			rwMutex.RLock()
			fmt.Printf("Incremented current counter: %v\n", counter)
			rwMutex.RUnlock()
			fmt.Println()

			rwMutex.RLocker()
		}()
	}

	wg.Wait()
	rwMutex.RLock()
	fmt.Println(counter)
	rwMutex.RUnlock()
}
