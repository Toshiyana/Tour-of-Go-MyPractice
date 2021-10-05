// フィボナッチ数列の実装
package main

import "fmt"

// fibonacci 数列を返す関数（クロージャ）
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		ret := a
		a, b = b, a+b
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
