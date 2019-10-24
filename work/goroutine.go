package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func main() {
	say("Hello")
	go say("Hello")

	go func(str string) {
		fmt.Println(str)
	}("Hello World")

	time.Sleep(time.Second)
	fmt.Println("Work Done")
}
