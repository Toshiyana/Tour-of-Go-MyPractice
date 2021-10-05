// 呼び出すメソッドを示す型がインターフェースのタプル内に存在しない場合、
// nilインタフェースのメソッドを呼び出すと、ランタイムエラーになる。

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
