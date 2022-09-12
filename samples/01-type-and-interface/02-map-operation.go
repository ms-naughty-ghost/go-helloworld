package main

import "fmt"

/*
マップのキーの順序は意図的にランダムにされる。追加や削除を行った順番も関係ないため
繰り返し処理を実行する場合も、実行ごとに順序付けが保証されない。
*/
func main() {
	// 空で初期化
	empty := map[string]int{}
	fmt.Printf("空マップ %v\n", empty)
	// make関数を使って初期化
	mapMake := make(map[string]int)
	fmt.Printf("makeで初期化 %v\n", mapMake)
	// 予め容量を確保した状態で初期化
	mapCap := make(map[string]int, 10)
	fmt.Printf("%v\n", mapCap)
	// 値およびキーの存在確認
	human := map[string]int{
		"John": 20,
		"Mike": 30,
	}
	fmt.Printf("%+v\n", human)
	age, ok := human["John"]
	fmt.Printf("値 %v, キーの有無 %v\n", age, ok)
	// マップに追加
	human["Marine"] = 31
	fmt.Printf("追加 %v\n", human)
	// キーと値の削除
	delete(human, "Mike")
	fmt.Printf("削除:%v\n", human)
	// キーのユニーク性を利用して配列の重複排除処理
	names := []string{"pekora", "aqua", "miko", "pekora", "suisei", "shion", "aqua"}
	// すべてユニークであることを考慮して、対象のスライスの長さ分の容量を確保する
	unique := make([]string, 0, len(names))
	// 存在有無チェックとしてマップを利用
	// 値を空の構造体にすることで、メモリのアロケーションをゼロに出来る
	m := make(map[string]struct{})
	for _, v := range names {
		if _, ok := m[v]; ok {
			continue
		}
		unique = append(unique, v)
		m[v] = struct{}{}
	}
	fmt.Printf("重複排除後 %v\n", unique)
}
