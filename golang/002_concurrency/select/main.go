package main


import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "every 500ms"
		time.Sleep(time.Millisecond * 500)
	}()

	go func() {
		c2 <- "Every two seconds"
		time.Sleep(time.Second * 2)
	}()

	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(c2)
		}
	}
	

}