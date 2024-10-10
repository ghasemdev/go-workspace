package main

import "fmt"

func main() {
	x := 10

	if x > 9 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is  not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	a := 11.0
	b := 20.0

	// frac = 0.55
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	switch_statements()
}

func switch_statements() {
	x := 2

	switch x {
	case 1:
		{
			fmt.Println("one")
		}
	case 2:
		{
			fmt.Println("two")
		}
	default:
		{
			fmt.Println("many")
		}
	}

	switch {
	case x > 1:
		{
			fmt.Println("one")
		}
	}
}
