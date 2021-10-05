// ポインタ：値のメモリアドレス。
// &i: 値iのポインタを取り出す
// *p: ポインタpの指す先の値を取り出す. *pとして値を取り出す場合、pはポインタである必要がある。

package main

import "fmt"

func main() {
	i := 1

	fmt.Println("&i:", &i) // アドレスを取り出す
	//fmt.Println(*i) // iは値でポインタではないので、*iはエラー

	fmt.Println("")

	k := &i                // iのポインタを代入
	fmt.Println("k:", k)   // ポインタを出力
	fmt.Println("*k:", *k) // 値を出力

	fmt.Println("")

	*k = 2 // 値の書き換え（*kとiのポインタは同じなので、iも書き換わる）
	fmt.Println("*k:", *k)
	fmt.Println("i:", i)
}
