package main

import "fmt"

type Trade struct {
	// Public
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price

	if trade.Buy {
		value = -value
	}

	return value
}

func main() {
	t1 := Trade{"APPLE", 10, 99.98, true}
	fmt.Printf("t1: %+v\n", t1)
	fmt.Println(t1.Symbol)
	fmt.Println(t1.Value())
}
