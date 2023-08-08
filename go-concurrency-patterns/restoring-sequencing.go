package main

import (
	"fmt"
	"time"
)

type Message struct {
	str   string
	block chan bool
}

func main() {
	ch := fanIn(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		msg1 := <-ch
		fmt.Println(msg1.str)

		msg2 := <-ch
		fmt.Println(msg2.str)

		fmt.Println(<-msg1.block) // reset channel, stop blocking
		// fmt.Println(<-msg2.block)
	}
}

// fanIn is itself a generator
func fanIn(ch1, ch2 <-chan Message) <-chan Message { // receives two read-only channels
	new_ch := make(chan Message)
	go func() {
		for {
			new_ch <- <-ch1
		}
	}() // launch two goroutine while loops to continuously pipe to new channel
	go func() {
		for {
			new_ch <- <-ch2
		}
	}()
	return new_ch
}

func generator(msg string) <-chan Message { // returns receive-only channel
	ch := make(chan Message)
	blockingStep := make(chan bool) // channel within channel to control exec, set false default
	go func() {                     // anonymous goroutine
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("%s %d", msg, i), blockingStep}
			time.Sleep(time.Second)
			blockingStep <- true // block by waiting for input
		}
	}()
	return ch
}

// func gen(msg string) <-chan Message {

// }
