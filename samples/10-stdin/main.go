package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	// 最初に読み取り行数の取得
	var n, _ = strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		// スペース区切りの入力
		var s = strings.Split(sc.Text(), " ")
		fmt.Println(fmt.Sprintf("hello = %v, world = %v", s[0], s[1]))
	}
}
