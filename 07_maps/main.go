package main

import "fmt"

func main() {
	bar := map[string]float64{
		"AMAZON":    1700.0,
		"GOOGLE":    1130.0,
		"MICROSOFT": 100.5,
	}
	fmt.Println(len(bar))
	fmt.Println(bar["AMAZON"])
	fmt.Println(bar["APPLE"])
	value, ok := bar["APPLE"]

	if !ok {
		fmt.Println("APPLE not found")
	} else {
		fmt.Println(value)
	}

	bar["APPLE"] = 350.0
	fmt.Println(bar)

	delete(bar, "AMAZON")
	fmt.Println(bar)

	fmt.Println()

	for key := range bar {
		fmt.Println(key)
	}

	fmt.Println()

	for key, value := range bar {
		fmt.Printf("%s - %.2f \n", key, value)
	}
}
