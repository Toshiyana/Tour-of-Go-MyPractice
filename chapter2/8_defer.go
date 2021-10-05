// defer（意味：持ち越す）
// deferへ渡した関数は、関数の実行を、呼び出し元の関数の終わり（return）まで遅らせる
// deferへ渡した関数の引数の評価はすぐに行われるが、実行するのは最後。

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("Ken") //渡した関数はstackされる。LIFO(last-in-fast-out)の順で実行される。

	fmt.Println("hello")
}
