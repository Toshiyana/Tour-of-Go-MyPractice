// range: forループで、スライスやmapを一つずつ反復処理する際に利用
// スライスのrange：反復毎に２つの変数index, value(indexの場所の要素のコピー)を返す

package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}
