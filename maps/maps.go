package maps

import (
	"fmt"
	"maps"
	"slices"
	"time"
)

type Product struct {
	ProductId uint   `json:"productId" gorm:"product_id"`
	Name      string `json:"name" gorm:"name"`
	Currency  string `json:"currrency" gorm:"currency"`
}
type Investment struct {
	InvestmentId uint    `json:"investmentId"`
	Amount       float64 `json:"amount"`
	CreateAt     int64   `json:"createdAt"`
	UpdatedAt    int64   `json:"updatedAT"`
	Product
}

func NewFeatures() {
	hashMap := make(map[int]string)
	hashMap[1] = "Michael"
	hashMap[2] = "Jordan"

	keys := slices.Sorted(maps.Keys(hashMap)) // Grab all the keys
	values := slices.Sorted(maps.Values(hashMap))
	fmt.Println(values)

	for _, key := range keys {
		fmt.Println(key)
	}
}

func LearnMaps() {
	fmt.Println("Learn about maps in Go")

	// var investments map[int]Investment
	investments := make(map[int]Investment) // NOTE: Zero value of map is nil.

	investments[1] = Investment{
		InvestmentId: 1,
		Amount:       2,
		CreateAt:     time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		Product: Product{
			ProductId: 1,
			Name:      "BTC Income Gilgamesh",
			Currency:  "BTC",
		},
	}

	investments[2] = Investment{
		InvestmentId: 2,
		Amount:       10,
		CreateAt:     time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		Product: Product{
			ProductId: 2,
			Name:      "ETH Income Gilgamesh",
			Currency:  "ETH",
		},
	}

	/*
		Iterating map using Range loop
		for key, value := range mapVariable {
			fmt.Println(key, value)
		}
	*/

	for investmentId, investment := range investments {
		fmt.Println(investmentId, investment.Amount, investment.Currency)
	}

	// Check numbers of items
	fmt.Printf("number of investments before deletion: %d\n", len(investments))

	// Deleting the first investment from the map
	delete(investments, 1)

	// Check if the values exists nor not.
	_, exists := investments[1]
	if !exists {
		fmt.Printf("investment with id %d was not found!\n", 1)
	}

	// Checking the size of the map
	fmt.Printf("number of investments after deletion: %d\n", len(investments))
}
