package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("sub sleep...")
		time.Sleep(time.Second * 2)
	}
}
func main() {
	go sub()
	for {
		fmt.Println("time sleep...")
		time.Sleep(time.Second * 1)
	}
}
