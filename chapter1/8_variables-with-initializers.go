package main

import "fmt"

// 初期値を与えている場合、型をつけなくても良い（型推論）
var i, j = 1, 2

func main() {
	var c, python, java = 1.0, true, "yes"
	fmt.Println(i, j, c, python, java)
}
