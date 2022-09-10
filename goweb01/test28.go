package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "waiting" }()

	msg := <-messages
	fmt.Println(msg)
}