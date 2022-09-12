package main

import "fmt"

func main() {
	var i int32 = 100
	var j int64
	// エラーになる
	// j = i
	j = int64(i)
	fmt.Println(j)

	// 文字列からバイトスライスにキャスト
	m := "hello world"
	bm := []byte(m)
	fmt.Println("バイト配列", bm)
	// バイトスライスから文字列にキャスト
	fmt.Println("文字列", string(bm))

	// 型アサーション
	greeting := interface{}("hello world")
	// 型アサーションを利用して文字列として値を受け取る
	hello := greeting.(string)
	fmt.Println(hello)
	// 型の変換結果を受け取る
	n, ok := greeting.([]byte)
	fmt.Println(n, ok)

	// 型スイッチ
	switch greeting.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("this is integer")
	case string:
		fmt.Println("this is string")
	case bool:
		fmt.Println("this is boolean")
	default:
		fmt.Println("this is unknown type", greeting)
	}

	// ユーザー定義型を型スイッチで分岐処理
	type ErrorNoSuchKey struct{ error }
	type ErrorBadRequest struct{ error }
	do := func() error {
		return &ErrorBadRequest{}
	}
	switch do().(type) {
	case *ErrorBadRequest:
		fmt.Println("error bad request")
	case *ErrorNoSuchKey:
		fmt.Println("error no such key")
	default:
		fmt.Println("unknown error")
	}
}
