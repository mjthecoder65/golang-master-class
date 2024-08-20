package time

import (
	"fmt"
	"sync"
	"time"
)

const NUMBER_OF_WORKERS int = 5

type Candle struct {
	ID        int64   `json:"id"`
	Timestamp int64   `json:"timesamp"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
}

func CandleWorker(wg *sync.WaitGroup, id int, receiver <-chan Candle) {
	defer wg.Done()

	for candle := range receiver {
		fmt.Printf("received candle: %v\n", candle)

		time.Sleep(2 * time.Second)
		fmt.Println("candle has been processed")
	}
}

func FetchCandles() ([]Candle, error) {
	candles := []Candle{}

	candle := Candle{
		ID:        1,
		Timestamp: 1723900000,
		Open:      2,
		Close:     3,
		High:      5,
		Low:       1.5,
		Volume:    2000,
	}

	candles = append(candles, candle)

	return candles, nil
}

func SendCandleMesages(candles []Candle, senderChan chan<- Candle) {
	for _, candle := range candles {
		senderChan <- candle
	}
}

func CandleMain() {

	ticker := time.NewTicker(2 * time.Second)
	messages := make(chan Candle, NUMBER_OF_WORKERS)

	// Spin up the worker to be read to receiver.
	var wg sync.WaitGroup

	for id := 0; id < NUMBER_OF_WORKERS; id++ {
		wg.Add(1)
		go CandleWorker(&wg, id, messages)
	}

	for currentTime := range ticker.C {
		fmt.Println(currentTime)
		candles, err := FetchCandles()
		if err != nil {
			fmt.Println("failed to fetch candles")
		} else {
			SendCandleMesages(candles, messages)
		}
	}

	wg.Wait()
}
