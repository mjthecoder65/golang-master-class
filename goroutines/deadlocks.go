package goroutines

import (
	"fmt"
	"sync"
)

func MasteringDeadLocks() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Goroutine 1 sending to ch1")
		ch1 <- 1
		fmt.Println("Goroutine 1 waiting to receive from ch2")
		val := <-ch2
		fmt.Println("Goroutine 1 received:", val)
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Goroutine 2 sending to ch2")
		ch2 <- 2
		fmt.Println("Goroutine 2 waiting to receive from ch1")
		val := <-ch1
		fmt.Println("Goroutine 2 received:", val)
	}()

	wg.Wait()
}
