package main

import (
	"fmt"

	mutex "github.com/nikodem-kirsz/go-engineering/concurency/mutexes/internal"
)

func main() {
	// testing
	fmt.Println("Running 3 concurrent go routines updating map key value pairs with channel mutex.")
	mutex.ChannelMutexTest()
	fmt.Println("Running 100 concurrent go routines updating counter with custom mutex implemantation")
	mutex.RWMutexTest()
}
