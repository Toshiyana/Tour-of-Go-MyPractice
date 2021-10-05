// 特徴的！
package main

import "fmt"

// 変数の後に型をつける
// 書き方1
func add(x int, y int) int {
	return x + y
}

// 書き方2 (2つ以上の変数が同じ型の場合、最後の型を残して省略可能)
func subtract(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(subtract(3, 5))
}
