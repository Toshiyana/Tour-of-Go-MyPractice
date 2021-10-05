package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// ifの条件式部分で、変数宣言が可能（変数は、if,elseのスコープ内でのみ有効）
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// %g: float64
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// Printlnもまとめて書くことが可能(pow2つが先に実行されてから、printlnが実行される)
	fmt.Println(
		pow(2, 3, 10),
		pow(2, 3, 7),
	)
}
