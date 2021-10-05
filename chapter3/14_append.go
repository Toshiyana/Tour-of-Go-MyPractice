// スライスへの要素の追加：append()(https://go-tour-jp.appspot.com/moretypes/15#:~:text=%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E3%81%AE%E8%A9%B3%E7%B4%B0%E3%81%AF-,documentation,-%E3%82%92%E5%8F%82%E7%85%A7%E3%81%97)
// 定義
// func append(s []T, vs ...T) []T
// s: 追加元のT型のスライス、vs: 追加するT型の変数群
// 戻り値：追加元のスライスと追加する変数群を合わせたスライス

// 元の配列sが変数群の追加の際に容量が小さい場合
// より大きいサイズの配列を割り当て直す(戻り値のスライスの割り当て先は新しくなる)

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
