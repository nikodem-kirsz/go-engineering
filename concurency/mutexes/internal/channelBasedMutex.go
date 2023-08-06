package mutexes

import (
	"fmt"
	"sync"
)

type Container struct {
	counters map[string]int
	mutex    chan struct{} // Channel to simulate mutex behaviour
}

func NewContainer() *Container {
	return &Container{
		counters: make(map[string]int),
		mutex:    make(chan struct{}, 1), // Buffered channel with capacity 1
	}
}

func (c *Container) inc(name string) {
	c.mutex <- struct{}{} // Acquire the mutex (block if another goroutine holds it)
	c.counters[name]++
	<-c.mutex // Rekease tge mutex, allowing to next goroutine to acquire it
}

func ChannelMutexTest() {
	c := NewContainer()

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)

	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("c", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
