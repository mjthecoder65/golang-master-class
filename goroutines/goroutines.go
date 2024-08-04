package goroutines

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

/*
	Goroutines allows us to write concurrent programs in Go.

	Goroutine is a lightweight thread of execution.

	Channels:
	Mutex:
	Select:
	Go runtime scheduler: Handle the distribution of work.
	Use different algorithms to balances loads amoung goroutines.

*/

const (
	DB_USER     = "db_user"
	DB_PASSWORD = "passowrd"
	DB_NAME     = "crypto_klines"
	SSL_MODE    = "disable"
)

type Candle struct {
	Symbol    string  `json:"symbol" db:"symbol"`
	OpenTime  string  `json:"open_time" db:"open_time"`
	Open      float64 `json:"open" db:"open"`
	High      float64 `json:"high" db:"high"`
	Low       float64 `json:"low" db:"low"`
	Close     float64 `json:"close" db:"close"`
	Volume    float64 `json:"volume" db:"volume"`
	CloseTime int64   `json:"close_time" db:"close_time"`
}

func Fetch(url string, channel chan<- string) {
	res, err := http.Get(url)

	if err != nil {
		channel <- fmt.Sprintf("Error fetchting %s : %v", url, err)
		return
	}

	defer res.Body.Close()

	channel <- fmt.Sprintf("Fetched %s", url)
}

func ReadFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	defer file.Close()

	// Reading the file here.
}

func LearnGoRoutines() {
	fmt.Println("Learning about Go routines")

	ch := make(chan string)
	urls := []string{"http://google.com"}

	var wg sync.WaitGroup // Used to wait for all goroutines to finish.

	for _, url := range urls {
		wg.Add(1) // Is called before starting the goroutine.
		go func(url string) {
			defer wg.Done() // defered function is called at the end of the surrounding function.
			Fetch(url, ch)
		}(url)
	}

	go func() {
		for message := range ch {
			fmt.Println(message)
		}
	}()

	wg.Wait()
	close(ch)
}
