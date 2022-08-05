package main

import (
	"fmt"
)

func main() {
	// This deadlocks bc normal channels require a receiver to exist ahead of time.
	// c := make(chan string)
	// c <- "hello"
	// fmt.Println(<-c)

	// Buffered channels can be used before a receiver has started listening
	// Just don't overflow the buffer
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
}