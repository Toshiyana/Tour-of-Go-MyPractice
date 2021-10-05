package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// ポインタレシーバを持つメソッド
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ポインタを引数に持つ関数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // (&v).Scale(10)と解釈
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	fmt.Println("")
	p := &Vertex{3, 4}
	p.Scale(10)
	fmt.Println(p)
	ScaleFunc(p, 10)
	fmt.Println(p)

	fmt.Println("")
	fmt.Println(p, *p, &p)
}
