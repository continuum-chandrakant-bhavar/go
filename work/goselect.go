package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "First"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Second"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Received 1", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}

	select {
	case msg1 := <-c2:
		fmt.Println("Received 1", msg1)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
