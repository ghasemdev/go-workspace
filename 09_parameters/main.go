package main

import "fmt"

func doubleAt(values []int, i int) { // pass by ref
	values[i] *= 2
}

func double(n int) { // pass by value
	n *= 2
}

func doublePtr(n *int) { // pass by value
	*n *= 2
}

func main() {
	values := []int{1, 2, 3}
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val)
	fmt.Println(val)

	doublePtr(&val)
	fmt.Println(val)
}
