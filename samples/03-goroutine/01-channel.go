package main

import (
	"fmt"
)

func main() {
	// 双方向チャネル
	var ch1 chan int
	// 受信専用チャネル
	// var recieverCh <-chan int
	// 送信用チャネル
	// var senderCh chan<- int
	ch1 = make(chan int, 10)
	ch2 := make(chan int)
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	// バッファサイズをこえるのでdeadlockされる
	// ch2 <- 0
	// データの送信
	ch1 <- 1
	fmt.Println(len(ch1))
	ch1 <- 2
	ch1 <- 3
	fmt.Println(len(ch1))
	// データ受信
	r := <-ch1
	fmt.Println("受信データ:", r)
	fmt.Println(len(ch1))
}
