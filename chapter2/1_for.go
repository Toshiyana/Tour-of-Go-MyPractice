// for文
// goでは、forの条件式を書く（）が必要ない

package main

import (
	"fmt"
)

func main() {

	// 書き方1
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	// 書き方2(sum1とは出力違う)
	// 初期化と後処理の記述はなくてもおけ
	// = C言語などのwhileは、Goではforで書ける
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
		fmt.Println(sum2) //2,4,8,16
	}

	// 無限ループ
	for {

	}
}
