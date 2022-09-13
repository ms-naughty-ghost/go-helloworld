package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("sample2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// 1行ずつ読み込み
	fr := bufio.NewScanner(f)

	for fr.Scan() {
		fmt.Println(fr.Text())
	}

	if err := fr.Err(); err != nil {
		return
	}
}
