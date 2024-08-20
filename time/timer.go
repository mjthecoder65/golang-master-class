package time

import (
	"fmt"
	"time"
)

/*
	We often want to execute Go code at some point
	in the future, or repeatedly at some interval.
	Go's built-in timer and ticker features make both
	of these tasks easier.

	Timer: Represent a single event in the future.
	You tell the timer, how long you want to wait, and it provides
	a channel that will notified at that time.

	time.Sleep(): Cannot be cancelled
	timer: can be stopped before it fires.
*/

func LearnAboutTimers() {
	timer1 := time.NewTimer(2 * time.Second) // Timer will wait for two seconds.

	// This blocks on the timer's channel C until it sends
	// a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second) // Wait for the goroutine to complete
}

// Handle
func Processor(deliver chan<- string, receiver <-chan string, id int) {
	// deliver chan will be used to send messages to other goroutines.
	// receiver channel will be used to only receive.
	// This specificity is used to make type safer programs.
}
