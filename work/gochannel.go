package main

import (
	"fmt"
	"time"
)

func working(done chan bool) {
	fmt.Println("Working .....")
	time.Sleep(time.Second)
	fmt.Println("Done ....")
	done <- true
}

func main() {
	message := make(chan string)
	go func() {
		message <- "Hello"
	}()

	msg := <-message
	fmt.Println(msg)

	done := make(chan bool)
	go working(done)
	// blocking worked to complet
	<-done
}
