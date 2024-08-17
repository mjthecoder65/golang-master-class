package goroutines

import (
	"fmt"
	"time"
)

/*
	Channels are the message pipes that connect
	goroutines in Go. They serve as a means for goroutines
	to communicate with each other.

	Sytanx:
		channel := make(chan <type>)
		Example:
		messages := make(chan string)

	Channels are typed by the values they convey.


	Sending and Receiving messages:
	Sending Sytanx:
		channel <- value

	Receiving Sytanx:
		value := <- channel

	Importantly:
		By default sends and receives block until both senders
		and receivers are ready.
		This property helps wait for messages without using any
		other synchronization methods.
*/

/* Channel Buffering

By default, channels are unbuffered, meaning that they will only
accept sends (chan <-) if there is a corresponding concurrent receive (<-chan)
ready to receive the value.

Buffered channels, accepts a limited
number of values without a correponding receiver for those values.
*/

func BufferedChannels() {
	const NUMBER_OF_MESSAGES int = 5
	messages := make(chan interface{}, NUMBER_OF_MESSAGES)

	messageList := [NUMBER_OF_MESSAGES]string{"Hello", "world", "Go", "Buffered", "Channel"}

	for _, message := range messageList {
		messages <- message
	}

	close(messages) // Remember to close the channel when using for ranged loop.

	for message := range messages {
		fmt.Println(message)
	}
}

func CandleWorker(done chan bool) {
	fmt.Println("Processing candles....")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
	done <- true
}

func ChannelSynchronization() {
	/* Here we are using blocking receive to wait for other other goroutine to finish.
	//However when waiting for multiple goroutines, It is recommended to use a WaitGroup.

	*/
	done := make(chan bool, 1)
	go CandleWorker(done)

	<-done // Using blocking receive. Blocks until receive notification from CandleWorker.
}

func Ping(pings chan<- string, message string) {
	pings <- message
}

func Pong(pings <-chan string, pongs chan<- string) {
	message := <-pings
	pongs <- message
}

func ChannelDirection() {
	/*
	 When using channels as function parameter,
	 you can specify if a channel is meant to only send or receive values.
	 This specificity increases the type-safety of the program.
	*/

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	Ping(pings, "passed message")
	Pong(pings, pongs)

	fmt.Println(<-pongs) // Using blocking receive here.
}

func Channels() {
	messages := make(chan string, 2)
	go func() {
		// Sending message to the channel
		messages <- "ping"
	}()

	messages <- "Hello" // Send messages but there is not receiver.

	// Using a blocking receive to wait for all other goroutines to finish.
	message := <-messages // main() goroutines will block until the messages is received.
	fmt.Println("Received: ", message)
}
