// 特徴的！
package main

import "fmt"

// 戻り値に名前をつけられる（ここではx, y）
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// 名前をつけた場合。returnには何も書かなくても値が返される（naked-return）
	// 用途：短い関数でのみnaked-returnを利用。長いと、可読性が下がる。
	return
}

func main() {
	fmt.Println(split(17))
}
