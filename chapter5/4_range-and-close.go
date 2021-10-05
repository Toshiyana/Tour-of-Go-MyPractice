// range and close
// チャネルのclose：送り手がこれ以上送信する値がないことを示すために、チャネルをcloseできる。

// チャネルがcloseされているかどうかをbool値で確認。
// 受診の指揮の２つ目のパラメータを割り当て、
// 受信する値がない and チャネルが閉じているなら、ok = false
//v, ok := <-ch

// 注意
// * 受け手でなく、送り手のチャネルだけcloseする（closeした受け手のチャネルに送信すると、エラーになる）
// * チャネルをcloseするのはこれ以上値が来ないことを受けてが知る必要がある時だけ（ファイルと異なり、通常はcloseする必要ない）。
//   例として、rangeループを終了する場合など。

package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //送り手のclose
}

func main() {
	c := make(chan int, 10) //チャネル幅10
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
