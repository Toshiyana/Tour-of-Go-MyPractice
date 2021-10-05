package main

import "fmt"

// 特徴的！
// 変数宣言(ただし、関数の中でのみ)：varの代わりに、:=が利用可能。（暗黙的型宣言 + 型推論）
// 関数の外ではvar必要

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
