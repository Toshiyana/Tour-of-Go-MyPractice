// methodは、structの型だけでなく（例：1_methods.go）、任意の型（type）でも宣言可能。
// レシーバを伴うmethodの宣言をする場合、レシーバ型が同じパッケージにある必要がある。

// 以下の例：自作の数値型のMyfloatで、Absメソッドを定義している。

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	// fはMyfloat型で、returnはfloat64型だから変換する必要あり
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
