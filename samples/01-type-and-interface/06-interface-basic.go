package main

import "fmt"

// 実装を隠蔽し、振る舞いのみを共通化する。
type CryThing interface {
	Cry() string
}

type Cat struct{}

// Cryメソッドを実装しているのでCryThingインターフェースを満たしている。
func (c Cat) Cry() string {
	return "Nya"
}

func main() {
	var cat CryThing = Cat{}
	fmt.Println(cat.Cry())
}
