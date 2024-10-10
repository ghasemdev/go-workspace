package main

import "fmt"

func add[T int | float64](a T, b T) T {
	return a + b
}

func divMod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	val2 := add(1.0, 2.3)
	fmt.Println(val2)

	div, mod := divMod(7, 2)
	fmt.Printf("div = %d, mod = %d\n", div, mod)
}
