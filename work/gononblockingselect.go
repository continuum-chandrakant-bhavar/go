package main

import "fmt"

func main() {

	message := make(chan string)
	signal := make(chan string)

	select {
	case msg := <-message:
		fmt.Println("Received", msg)
	default:
		fmt.Println("No Message Received")

	}

	msg := "hi"

	select {
	case message <- msg:
		fmt.Println("Message Sent", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("Message Received......", msg)
	case sig := <-signal:
		fmt.Println("Signal Received....", sig)
	default:
		fmt.Println("No Activity....")

	}
}
