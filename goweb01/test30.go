package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
	fmt.Println("开始............")
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	fmt.Println("结束......",<-done)
}