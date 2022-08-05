package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	c := make(chan string)
	go count("cats", c)

	// for {
	// 	msg, open := <- c
	// 	fmt.Println(msg)

	// 	if !open {
	// 		break
	// 	}
	// }

	for msg := range c {
		fmt.Println(msg)
	}
	
}

func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // Only channel senders should close the channel
}