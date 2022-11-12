package main

import "fmt"

// FizzBuzz
// 1から整数を数える
// 3で割り切れる = Fizz
// 5で割り切れる = Buzz
// 3でも5でも割り切れる = FizzBuzz
func main() {
	n := 100
	for i := 1; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
