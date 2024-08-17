package time

import (
	"fmt"
	"time"
)

/*
	Select lets wait on multiple channel operations.
	Combining goroutine and channels with select isa powerful
	features in Go.

*/

func LearnSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("First goroutine finished")
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Second goroutine finished")
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-c1:
			fmt.Println("Received: ", message1)
		case message2 := <-c2:
			fmt.Println("Received: ", message2)
		}
	}
}
