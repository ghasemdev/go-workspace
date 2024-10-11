package main

import (
	"fmt"
)

func saveValues(values []int, index int) int {
	defer func() { // try
		if err := recover(); err != nil { // catch
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return values[index]
}

func main() {
	// array := []int{1, 2, 3, 4}
	// v := array[10]
	// fmt.Println(v)

	// file, err := os.Open("no-file")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// fmt.Println(file)

	v := saveValues([]int{1, 2, 3}, 5)
	fmt.Println(v)
}
