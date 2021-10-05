// チャネル：並列処理のgoroutine間を接続するパイプのイメージ。

// チャネル型：<-で値の送受信が可能な通り道。
// チャネルの生成
// ch := make(chan int)
// チャネルの送受信
// ch <- v    // v をチャネル ch へ送信する
// v := <-ch  // ch から受信した変数を v へ割り当てる

// 以下例
// スライスを２つに分割し、数値の総和を算出する処理を、2つのgoroutine間で作業を分担する。
// 両方のgoroutineの計算が終了すると、最終結果が算出される。

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // チャネルの生成
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
