package main

import "fmt"

// 変数のコピーコストを考えると参照渡しがいい
func Double(i int) {
	i = i * 2
}
func DoublePointer(i *int) {
	// アスタリスクで値参照
	*i = *i * 2
}

func main() {
	var i int = 100
	fmt.Println(i)
	fmt.Println(&i)

	// 参照渡し
	DoublePointer(&i)
	fmt.Println(i)
}
