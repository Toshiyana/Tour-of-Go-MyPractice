// 型変換
package main

import (
	"fmt"
	"math"
)

func main() {
	// 例1
	var i int = 42
	var t float64 = float64(i)
	var u uint = uint(t)

	// 例2（よりシンプルに書く場合）
	a := 42
	b := float64(a)
	c := uint(b)

	fmt.Println(i, t, u)
	fmt.Println(a, b, c)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
