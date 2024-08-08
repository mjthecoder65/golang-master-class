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

type Bid struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type Ask struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}
type OrderBook struct {
	Symbol        string `db:"symbol" json:"symbol"`
	Bids          []Bid  `db:"bids" json:"bids"`
	Asks          []Ask  `db:"asks" json:"asks"`
	LastUpdatedID int64  `db:"last_updated" json:"lastUpdatedId"`
}

func FetchOrderBooks(symbol string, limit int, client *binance.Client) OrderBook {
	/*
	 Depth: Refers to the market depth, which is the measure of supply and demand
	 for a particular trading pair (such as BTCUSDT) var varies price level.

	 It shows the available orders in the order book, and is crucial to understanding
	 market liquidity and potential price movement.
	*/

	depthService := client.NewDepthService().Symbol(symbol).Limit(1)

	depth, err := depthService.Do(context.Background())

	if err != nil {
		fmt.Printf("error: failed to fetch order books for %s\n", symbol)
	}

	var orderBook OrderBook
	orderBook.LastUpdatedID = depth.LastUpdateID

	for _, ask := range depth.Asks {
		orderBook.Asks = append(orderBook.Asks, Ask{
			Price:    ask.Price,
			Quantity: ask.Quantity,
		})
	}

	for _, bid := range depth.Bids {
		orderBook.Bids = append(orderBook.Bids, Bid{
			Price:    bid.Price,
			Quantity: bid.Quantity,
		})
	}

	return orderBook
}

func FetchCandleEveryTwoMinutes(symbol string, interval string, limit int, client *binance.Client) {
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

	go FetchCandleEveryTwoMinutes(symbol, interval, limit, client)

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
