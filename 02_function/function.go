package main

import "fmt"

func main() {
	var x int
	var y int

	x = 5
	y = 10

	fmt.Printf("x = %v, type of %T \n", x, x)
	fmt.Printf("x = %v, type of %T \n", y, y)

	mean := (float64(x) + float64(y)) / 2.0
	fmt.Printf("result: %v, type of %T", mean, mean)
}
