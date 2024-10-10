package main

import "fmt"

func main() {
	book := "Hello My Boy!!"
	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v, (type) %T \n", book[0], book[0])

	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:4])
	fmt.Println(book[:4] + "hi")
	poem := `
	Hello;
	WW
	`
	fmt.Println(poem)
}
