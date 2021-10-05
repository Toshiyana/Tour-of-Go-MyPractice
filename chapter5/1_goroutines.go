// goroutine: Goのランタイムに管理される軽量なスレッド。（並列処理）
// goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要あり。

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// この処理がないと、先にmain関数の処理が終わってしまい、goroutineの関数が実行されない
		time.Sleep(100 * time.Millisecond)

		fmt.Println(s)
	}
}

func main() {
	// 並列処理ない場合
	// say("world")
	// say("hello")

	// 並列処理ある場合
	go say("world") // goroutine
	say("hello")
}
