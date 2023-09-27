package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// c := boring("boring!")
	c := fanIn(boring("joe"), boring("ann"))

	for i := 0; i < 10; i++ {
		// fmt.Printf("You say: %q\n", <-c)
		message := <-c
		fmt.Println(message.str)
		message.waitForIt <- true
	}
	fmt.Printf("You are boring: I'm leaving")
}

type Message struct {
	str       string
	waitForIt chan bool
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	// go func() {
	// 	for {
	// 		c <- <-input1
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		c <- <-input2
	// 	}
	// }()
	go func() {
		defer close(c)
		for {
			select {
			case msg1, ok := <-input1:
				if !ok {
					input1 = nil
					continue
				}
				c <- msg1
			case msg2, ok := <-input2:
				if !ok {
					input2 = nil
					continue
				}
				c <- msg2
			}
		}
	}()
	return c
}

func boring(msg string) <-chan Message {
	waitForIt := make(chan bool)
	c := make(chan Message)

	go func() { // We launch a go routine from inside a function
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c // Return the channel to
}
