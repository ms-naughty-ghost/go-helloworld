package main

import "fmt"

// https://qiita.com/rock619/items/14eb2b32f189514b5c3c
func main() {
	// すべての型に使える%v
	fmt.Printf("%v,%v\n", 100, true)
	// フィールド名を出力する%+v
	var man struct {
		Name string
		Age  int
	}
	man.Name = "Mike"
	man.Age = 30
	fmt.Printf("%+v\n", man)
}
