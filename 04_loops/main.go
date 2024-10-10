package main

import "fmt"

func main() {
	count := 3
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	fmt.Println("+++++++++++++++++++")

	for i := 0; i < count; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("+++++++++++++++++++")

	for i := 0; i < count; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("+++++++++++++++++++")

	a := 0
	for a < count {
		fmt.Println(a)
		a++
	}

	fmt.Println("+++++++++++++++++++")

	b := 0
	for {
		fmt.Println(b)
		if b > 3 {
			break
		}
		b++
	}
}
