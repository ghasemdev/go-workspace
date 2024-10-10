package main

import "fmt"

type Trade struct {
	// Public
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"APPLE", 10, 99.98, true}
	fmt.Printf("t1: %+v\n", t1)
	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MICROSOFT",
		Volume: 15,
		Price:  99.98,
		Buy:    false}
	fmt.Printf("t2: %+v\n", t2)
}
