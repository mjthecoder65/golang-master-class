package time

import (
	"fmt"
	"time"
)

/*
	Timers are for when you want to do something once in the future.
	tickers, are for when you want do something repeteatedly at regular intervals.

	Select: You can think of select as switch except that it works for a channel.
	// It is used to wait on multiple channels. And one action is performed
	on goroutine, when a values is send to either of the channels.

*/

func Tickers() {
	fmt.Println("Learning about tickers")

	// this creates a ticker, that send current time to a
	// after every half a second.
	ticker := time.NewTicker(500 * time.Millisecond)

	// This channel will be used to signal the goroutine to stop.
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Sleep for 1.6 seconds to give enough time for tickers to be send to the channel
	time.Sleep(10 * time.Second)

	// Stop sending ticks
	ticker.Stop()

	// Signal the goroutine to stop running.
	done <- true
	fmt.Println("ticker has been stopped")
}
