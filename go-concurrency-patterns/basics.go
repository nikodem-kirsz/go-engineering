package main

import (
	"fmt"
	"time"
)

func main() {
	go regular_print("Hello")
	fmt.Println("Second print statement!")
	time.Sleep(3 * time.Second)
	fmt.Println("Third print statement!") // when main returns, the goroutines also end
}

func regular_print(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
