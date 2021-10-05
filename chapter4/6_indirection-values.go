package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 値レシーバを持つメソッド
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 値を引数に持つ関数
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	fmt.Println("")
	p := &Vertex{4, 3} // ポインタを代入
	fmt.Println(p)
	fmt.Println(p.Abs())     // (*p).Abs()として解釈
	fmt.Println(AbsFunc(*p)) //値*pを引数とする
}
