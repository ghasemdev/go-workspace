package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	}

	return math.Sqrt(n), nil
}

func main() {
	result, err := sqrt(4.0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result2, err2 := sqrt(-2.0)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(result2)
	}
}
