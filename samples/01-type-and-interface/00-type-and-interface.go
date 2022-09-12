package main

import "fmt"

func main() {
	// 符号付き整数
	// int, uintは32bit実装,64bit実装で変化する
	var i int = 9223372036854775807
	fmt.Println("int", i)
	var i8 int8 = 127 // -128 ～ 127
	fmt.Println("int8", i8)
	var i16 int16 = 32767 // -32768 ～ 32767
	fmt.Println("int16", i16)
	var i32 int32 = 217483647 // -2147483648 ～ 2147483647
	fmt.Println("int32", i32)
	var i64 int64 = 9223372036854775807 // -9223372036854775808 ～ 9223372036854775807
	fmt.Println("int64", i64)
	var r rune = 2147483647 // -2147483648 ～ 2147483647
	fmt.Println("rune", r)
	// 符号なし整数
	var ui uint = 0
	// var ui uint = -1
	fmt.Println("uint", ui)
	var ui8 uint8 = 255 // 0 ～ 255
	fmt.Println("uint8", ui8)
	var ui16 uint16 = 65535 // 0 ～ 65535
	fmt.Println("uint16", ui16)
	var ui32 uint32 = 4294967295 // 0 ～ 4294967295
	fmt.Println("uint32", ui32)
	var ui64 uint64 = 18446744073709551615 // 0 ～ 18446744073709551615
	fmt.Println("uint64", ui64)
	var uip uintptr = 0xc0000ba000 // ポインタの値をそのまま格納するのに十分な大きさの符号なし整数。
	fmt.Println("uintptr", uip)
	var b byte = 255 // 0 ～ 255
	fmt.Println(b)
	// 浮動小数点数
	var f32 float32 = 0.000000000000001
	fmt.Println("float32", f32)
	var f64 float64 = 0.00000000000000001
	fmt.Println("float64", f64)
	// 複素数
	var comp64 complex64
	fmt.Println(comp64)
	var comp128 complex128
	fmt.Println(comp128)
	// 文字列
	var str string = "hello world"
	fmt.Println("string", str)
	// 真偽値
	var bool_value bool
	fmt.Println("bool", bool_value)
	// エラー
	var e error
	fmt.Println("error", e)

	// 構造体の宣言
	type human struct {
		Name string
		Age  int
	}
	// 構造体の定義
	var man human = human{
		Name: "ghost",
		Age:  30,
	}
	fmt.Println("human struct", man)

	// 配列
	// ゼロ値で初期化
	var array [5]int
	fmt.Println("ゼロ値で初期化", array)
	// 5つの要素を初期化
	arrayLiteral := [5]int{0, 1, 2, 3, 4}
	fmt.Println("5つの要素を初期化", arrayLiteral)
	// 要素数から配列サイズを推定
	arrayInference := [...]int{0, 1, 2, 3, 4}
	fmt.Println("要素数から配列サイズを推定", arrayInference)
	// 配列のインデックスと値を指定
	arrayIndex := [...]int{2: 5, 4: 1, 7: 8}
	fmt.Println("配列のインデックスと値を指定", arrayIndex)

	// スライス
	var slice []int
	slice = append(slice, 2, 5) // データの増減可能
	fmt.Println("スライス", slice)

	// マップ
	// ゼロ値で初期化
	var m = map[string]int{
		"John": 32,
		"Mike": 40,
	}
	fmt.Println("map", m)
}
