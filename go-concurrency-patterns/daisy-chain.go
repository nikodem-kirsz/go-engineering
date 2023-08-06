package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 1000

	var channels [n+1]chan int

	for i := range channels {
		go f(channels[i], channels[i+1])
	}

	go func(c chan<- int) { 
		c <- 1 
	} (channels[n])

	fmt.Println(<-channels[0])
}