package goroutines

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
)

const (
	BINANCE_API_KEY                 string = "None" // Load from .env
	BINANCE_API_SECRET_KEY          string = "None" // Load from .env
	NUMBER_OF_WORKERS               int    = 5
	CANDLE_CHANNEL_BUFFER_COUNT     int    = 1000
	ORDER_BOOK_CHANNEL_BUFFER_COUNT int    = 1000
)

var (
	candleChannel    = make(chan Candle, CANDLE_CHANNEL_BUFFER_COUNT)
	orderBookChannel = make(chan OrderBook, ORDER_BOOK_CHANNEL_BUFFER_COUNT)
	wg               sync.WaitGroup
)

type Candle struct {
	Symbol    string `db:"symbol" json:"symbol"`
	OpenTime  int64  `db:"open_time" json:"open_time"`
	CloseTime int64  `db:"close_time" json:"close_time"`
	Open      string `db:"open" json:"open"`
	Close     string `db:"close" json:"close"`
	High      string `db:"high" json:"high"`
	Low       string `db:"low" json:"low"`
	Volume    string `db:"volume" json:"volume"`
}

type OrderBook struct {
	Symbol        string     `db:"symbol" json:"symbol"`
	Bids          [][]string `db:"bids" json:"bids"`
	Asks          [][]string `db:"asks" json:"asks"`
	LastUpdatedID int        `db:"last_updated" json:"lastUpdatedId"`
}

func FetchCandles(symbol string, interval string, limit int, client *binance.Client) {
	klineService := client.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit)
	klines, err := klineService.Do(context.Background())

	if err != nil {
		return
	}

	candles := []Candle{}

	for _, kline := range klines {
		candle := Candle{
			Symbol:    symbol,
			OpenTime:  kline.OpenTime,
			CloseTime: kline.CloseTime,
			Open:      kline.Open,
			Close:     kline.Close,
			High:      kline.High,
			Low:       kline.Low,
			Volume:    kline.Volume,
		}
		candles = append(candles, candle)
	}

	for _, candle := range candles {
		candleChannel <- candle
	}
}

func fetchCandleEveryTwoMinutes(symbol string, interval string, limit int, client *binance.Client) {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		FetchCandles(symbol, interval, limit, client)
	}
}

func CollectCandles() {
	client := binance.NewClient(BINANCE_API_KEY, BINANCE_API_SECRET_KEY)

	var interval string = "1m"
	var limit int = 1000
	var symbol string = "BTCUSDT"

	go fetchCandleEveryTwoMinutes(symbol, interval, limit, client)

	// Start worker goroutines
	for i := 0; i < NUMBER_OF_WORKERS; i++ {
		wg.Add(1)
		go ProcessCandles()
	}
	wg.Wait()
}

func ProcessCandles() {
	defer wg.Done()
	for candle := range candleChannel {
		fmt.Printf("INFO: Received candle for %s ", candle.Symbol)
		fmt.Println(candle)
	}
}

func ProcessOrderBook() {
	defer wg.Done()

	for orderBook := range orderBookChannel {
		fmt.Printf("INFO: Received candle for %s ", orderBook.Symbol)
	}

}
