package main

import "fmt"

type Holomen struct {
	Name string
}

// (e *Holomen)はレシーバと呼ばれる。レシーバと紐つけられた関数をメソッドと呼ぶ。
// レシーバはメソッド実行時に引数のように扱われる。
// データに強く関連する処理を宣言する。
func (h *Holomen) ShowName() {
	fmt.Printf("タレント %v\n", h.Name)
}

func (h Holomen) Change(v string) {
	h.Name = v
}

func main() {
	pekora := &Holomen{
		Name: "兎田ぺこら",
	}
	pekora.ShowName()
	pekora.Change("星街すいせい") // レシーバが型なので値の更新は反映されない
	pekora.ShowName()
}
