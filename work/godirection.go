package main

import (
	"fmt"
)

//This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, message string) {
	pings <- message
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, "Passed Message")
	go pong(pings, pongs)
	fmt.Println(<-pongs)

}
