package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	// This will bock because nothing wa pushed to the channel
	// <-ch
	// fmt.Println("HERE")

	go func() {
		// Send number of the channel
		channel <- 3
	}()

	// Receive number from the channel
	value := <-channel
	fmt.Printf("got %d\n", value)

	// Send Multiple
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("Sending %d\n", i)
			channel <- i
			time.Sleep(time.Second)
		}
	}()

	// Receive Multiple
	for i := 0; i < 4; i++ {
		value := <-channel
		fmt.Printf("Received %d\n", value)
	}

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("Sending %d\n", i)
			channel <- i
			time.Sleep(time.Second)
		}
		close(channel)
	}()

	for i := range channel {
		fmt.Printf("Received %d\n", i)
	}
}
