package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	//生产者
	// ticker := time.NewTicker(1*time.Second)
	go func() {
		for i := 0; ; i++ {
			ch <- i
			//	time.Sleep(1 * time.Second)
		}
	}()

	for {
		fmt.Println("chan message is: ", <-ch)
		time.Sleep(1 * time.Second)
	}
	//time.Sleep(1 * time.Minute)
}
