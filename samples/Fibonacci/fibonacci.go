package main

import "fmt"

// フィボナッチ数 - シンプル再帰
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	for i := 0; i < 30; i++ {
		fmt.Println(fibonacci(i))
	}
}
