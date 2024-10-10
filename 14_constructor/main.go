package main

import (
	"fmt"
	"os"
)

type Trade struct {
	// Public
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func NewTrade(Symbol string, Volume int, Price float64, Buy bool) *Trade {
	if Symbol == "" {
		return nil
	}
	if Volume < 0 {
		return nil
	}
	if Price < 0 {
		return nil
	}

	trade := &Trade{
		Symbol: Symbol,
		Volume: Volume,
		Price:  Price,
		Buy:    Buy,
	}
	return trade
}

func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price

	if trade.Buy {
		value = -value
	}

	return value
}

func main() {
	t1 := NewTrade("APPLE", 10, 99.98, true)
	if t1 == nil {
		fmt.Println("Can't create trade")
		os.Exit(1)
	}
	fmt.Printf("t1: %+v\n", t1)
	fmt.Println(t1.Symbol)
	fmt.Println(t1.Value())
}
