// 練習問題：平方根を求める関数の実装

// メモ: アルゴリズムの詳細について興味がある人のために説明すると、
// 上の z² − xという式は、z² が最終的な期待値 x からどのくらい離れているかを表しています。
// 除算の 2z は z² の導関数で、z² の変化の大きさに応じて z の調整値を変化させます。 この一般的なアプローチはニュートン法と呼ばれています。
// 多くの関数で有効ですが、平方根で特に有効です。

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	// ニュートン法
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(z)
	}
	return z
}

func main() {
	// 自作Sqrt
	fmt.Println(Sqrt(2))
	// 標準Sqrt
	fmt.Println(math.Sqrt(2))
}
