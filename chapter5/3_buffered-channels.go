// バッファ：一時的に記憶する場所。チャネルはバッファとして扱える。

// バッファを持つチャネルの初期化:makeの2つ目の引数にバッファ長を与える。
// ch := make(chan int, 100)

// * バッファが詰まった時、チャネルへの送信をブロック
// * バッファが空の時、チャネルの受信をブロック

package main

import "fmt"

func main() {
	ch := make(chan int, 2) //バッファ長2
	ch <- 1
	ch <- 2
	//ch <- 3 //バッファ長2の場合、エラーになる(3ならok)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
