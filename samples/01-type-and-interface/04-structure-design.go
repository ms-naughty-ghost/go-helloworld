package main

// 先頭大文字
// パッケージ外からもアクセスできる
type Enterprise struct {
	Name    string
	Address string
}

// 先頭小文字
// パッケージ外からアクセスできない
type employee struct {
	Name string // パッケージ外からアクセスできる
	age  int    // パッケージ外からアクセスできない
}
