// Pointer receivers

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 変数レシーバを持つメソッド
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバを持つメソッド
// ポインタレシーバ*Vertexを持つScaleメソッドを定義。
// 特徴：レシーバが指す変数を変更可能。
// ポインタのない変数レシーバの場合、元のVertex変数のコピーを操作。
// ポインタのあるポインタレシーバの場合、元のVertex変数自体を操作。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
