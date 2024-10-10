package main

import "fmt"

func main() {
	foo := []string{"amir", "azimi", "parsclick"}
	fmt.Printf("foo = %v (type %T)\n", foo, foo)
	fmt.Println(len(foo))
	fmt.Println(foo[1])
	fmt.Println(foo[:2])

	fmt.Println("++++++++++++")

	for i := 0; i < len(foo); i++ {
		fmt.Println(foo[i])
	}

	fmt.Println("++++++++++++")

	for i, v := range foo {
		fmt.Printf("index %d, value %s \n", i, v)
	}
}
