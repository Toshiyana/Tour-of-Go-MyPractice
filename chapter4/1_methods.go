// method (Goにはclassがない)
// structによる型がclassの代わりのようなもので、型にmethodを定義できる。

// methodはreceiver引数を関数に取る（funcとmethod名の間に、receiver引数を記述）
// receiver引数には、定義した型の引数が入る

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method(レシーバ引数を伴う関数)
// 以下の例の場合、Vertexという型の要素に、Abs1()という関数を定義している。
// (v Vertex): レシーバ引数, レシーバ引数の型
// Abs1(): メソッド名
// float64：戻り値の方
func (v Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function（レシーバ引数を伴わないmethod = 通常の関数）
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	a := Vertex{3, 4}
	// method
	fmt.Println(a.Abs1())
	// function
	fmt.Println(Abs2(a))
}
