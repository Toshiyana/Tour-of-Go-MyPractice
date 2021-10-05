// Select文：goroutineを複数の通信操作で待たせる。
// selectは、複数あるcaseのいずれかが準備できるまでブロックし、準備ができたcaseを実行する。
// もし、複数のcaseの準備ができている場合、caseはランダムに選択される。

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
