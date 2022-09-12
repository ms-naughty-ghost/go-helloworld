package main

import "fmt"

// スライス
// https://github.com/golang/go/wiki/SliceTricks
func main() {
	sl := []int{10, 20, 30}
	fmt.Printf("slice %v, 長さ %v, 容量 %v\n", sl, len(sl), cap(sl))
	// スライスは要素数可変
	sl = append(sl, 50, 51, 52)
	fmt.Printf("slice %v, 長さ %v, 容量 %v\n", sl, len(sl), cap(sl))

	// makeを使った初期化
	// 長さと容量を指定
	sl2 := make([]int, 2, 3)
	fmt.Printf("slice2 %v, 長さ %v, 容量 %v\n", sl2, len(sl2), cap(sl2))
	// スライスの複製
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("dst %v, 長さ %v, 容量 %v\n", dst, len(dst), cap(dst))
	// スライス同士の結合
	src1, src2 := []int{1, 2}, []int{3, 4, 5}
	dst1 := append(src1, src2...)
	fmt.Printf("dst1 %v\n", dst1)
	// appendでスライド内の要素を削除する
	target := []int{1, 2, 3, 4, 5}
	delIndex := 3
	dst2 := append(target[:delIndex], target[delIndex+1:]...)
	fmt.Printf("削除後 dst2 %v\n", dst2)
	// copyでスライド内の要素を削除する
	dst3 := target[:delIndex+copy(target[delIndex:], target[delIndex+1:])]
	fmt.Printf("削除後 dst3 %v\n", dst3)
	// スライスの逆順並び替え
	sortTarget := []int{1, 2, 3, 4, 5}
	for i := len(sortTarget)/2 - 1; i >= 0; i-- {
		tmp := len(sortTarget) - 1 - i
		sortTarget[i], sortTarget[tmp] = sortTarget[tmp], sortTarget[i]
	}
	fmt.Printf("ソート %v\n", sortTarget)
}
