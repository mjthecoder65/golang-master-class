package goroutines

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	BINANCE_API_KEY                 string = "None" // Load from .env
	BINANCE_API_SECRET_KEY          string = "None" // Load from .env
	NUMBER_OF_WORKERS               int    = 5
	CANDLE_CHANNEL_BUFFER_COUNT     int    = 1000
	ORDER_BOOK_CHANNEL_BUFFER_COUNT int    = 1000
	DATABASE_DSN                    string = "host=localhost user=admin password=y7jHf&DNWG15 dbname=main port=5030 sslmode=disable"
)

var (
	candleChannel    = make(chan Candle, CANDLE_CHANNEL_BUFFER_COUNT)
	orderBookChannel = make(chan OrderBook, ORDER_BOOK_CHANNEL_BUFFER_COUNT)
	wg               sync.WaitGroup
)

type Assets struct {
	ID   uint   `db:"id" json:"id" gorm:"primary_key"`
	Name string `db:"name" json:"name"`
}

type Candle struct {
	ID        uint    `db:"id" json:"id" gorm:"primary_key"`
	Symbol    string  `db:"symbol" json:"symbol"`
	Interval  string  `db:"symbol" json:"interval"`
	OpenTime  int64   `db:"open_time" json:"open_time"`
	CloseTime int64   `db:"close_time" json:"close_time"`
	Open      float64 `db:"open" json:"open"`
	Close     float64 `db:"close" json:"close"`
	High      float64 `db:"high" json:"high"`
	Low       float64 `db:"low" json:"low"`
	Volume    float64 `db:"volume" json:"volume"`
}

func FetchCandles(symbol string, interval string, limit int, client *binance.Client) {
	klineService := client.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit)
	klines, err := klineService.Do(context.Background())

	if err != nil {
		return
	}

	candles := []Candle{}

	for _, kline := range klines {
		open, _ := strconv.ParseFloat(kline.Open, 64)
		close, _ := strconv.ParseFloat(kline.Close, 64)
		high, _ := strconv.ParseFloat(kline.High, 64)
		low, _ := strconv.ParseFloat(kline.Low, 64)
		volume, _ := strconv.ParseFloat(kline.Volume, 64)

		candle := Candle{
			Symbol:    symbol,
			Interval:  interval,
			OpenTime:  kline.OpenTime,
			CloseTime: kline.CloseTime,
			Open:      open,
			Close:     close,
			High:      high,
			Low:       low,
			Volume:    volume,
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
	 for a particular trading pair (such as BTCUSDT).

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
	ticker := time.NewTicker(time.Minute / 5)
	defer ticker.Stop()
	for range ticker.C {
		FetchCandles(symbol, interval, limit, client)
	}
}

func CollectCandlesMain() {
	db, err := gorm.Open(postgres.Open(DATABASE_DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database")
	} else {
		fmt.Println("Connected to the database...")
	}
	db.AutoMigrate(&Candle{})

	client := binance.NewClient(BINANCE_API_KEY, BINANCE_API_SECRET_KEY)

	var interval string = "1m"
	var limit int = 1000
	var symbol string = "BTCUSDT"

	go FetchCandleEveryTwoMinutes(symbol, interval, limit, client)

	for i := 0; i < NUMBER_OF_WORKERS; i++ {
		wg.Add(1)
		go ProcessCandles(db)
	}

	wg.Wait()
}

func SaveCandle(candle Candle, db *gorm.DB) {
	result := db.Create(&candle)

	if result.Error != nil {
		log.Printf("failed to save data to the database: %v", result.Error)
	} else {
		fmt.Println("INFO: saved candled the database")
	}
}

func AddOrUpdate(candle Candle, db *gorm.DB) {
	var existingCandle Candle
	result := db.Where("symbol = ? AND interval = ? AND open_time = ? AND close_time = ?",
		candle.Symbol, candle.Interval, candle.OpenTime, candle.CloseTime).
		First(&existingCandle)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Printf("Error Checking for existing candles: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		if err := db.Create(&candle).Error; err != nil {
			log.Printf("failed to create new candle: %v", err)
		} else {
			log.Println("New candle added.")
		}
	} else {
		existingCandle.Open = candle.Open
		existingCandle.High = candle.High
		existingCandle.Low = candle.Low
		existingCandle.Close = candle.Close
		existingCandle.Volume = candle.Volume
		existingCandle.CloseTime = candle.CloseTime

		if err := db.Save(&existingCandle).Error; err != nil {
			log.Printf("Error updating candle: %v", err)
		} else {
			log.Println("Candle have been updated.s")
		}
	}
}

func FindCandles(symbol string, interval string, page int, limit int, db *gorm.DB) []Candle {
	var candles []Candle

	offset := (page - 1) * limit

	result := db.Where("symbol = ? AND interval = ?", symbol, interval).
		Order("open_time DESC").
		Offset(offset).
		Limit(limit).
		Find(&candles)

	if result.Error != nil {
		log.Printf("There are not candles for provided search parameters")
	}

	return candles
}

func ProcessCandles(db *gorm.DB) {
	defer wg.Done()
	for candle := range candleChannel {
		SaveCandle(candle, db)
	}
}

// My first projects and that can used for backtesting.
func ProcessOrderBook() {
	defer wg.Done()

	for orderBook := range orderBookChannel {
		fmt.Printf("INFO: Received candle for %s ", orderBook.Symbol)
	}
}
