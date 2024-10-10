package main

import "fmt"

func cleanUp(name string) {
	fmt.Printf("Clean Up %s\n", name)
}

func worker() {
	defer cleanUp("A")

	fmt.Println("worker")
}

func main() {
	worker()
}
