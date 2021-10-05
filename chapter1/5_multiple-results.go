package main

import "fmt"

// 複数の戻り値を返すことも可能
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("world", "hello")
	fmt.Println(a, b)
}
