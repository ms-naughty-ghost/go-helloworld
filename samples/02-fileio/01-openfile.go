package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// 読み込みバッファーサイズ
	const bufferSize = 256
	var content []byte
	buffer := make([]byte, bufferSize) // 読み取りバッファー
	// ファイル終端まで読み込み
	for {
		n, err := f.Read(buffer)
		if 0 < n {
			// 読み込み処理
			content = append(content, buffer...)
		}
		if err == io.EOF {
			break // 終端
		}
		if err != nil {
			return // エラー
		}
	}

	fmt.Println(string(content))
}
