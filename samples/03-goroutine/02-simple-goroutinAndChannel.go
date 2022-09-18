package main

import (
	"fmt"
	"time"
)

func reciever(c chan int) {
	for {
		fmt.Println(<-c)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go reciever(ch1)
	go reciever(ch2)

	for i := 0; i < 100; i++ {
		ch1 <- i
		ch2 <- i
		fmt.Println("sleep...")
		time.Sleep(time.Millisecond * 50)
	}
	close(ch1)
	close(ch2)
}
