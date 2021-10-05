// Numeric Constants (数値の定数)
package main

import "fmt"

const (
	Big   = 1 << 100  // 左に100シフト
	Small = Big >> 99 // 右に99シフト
)

// 定数は状況によって、必要な型をとる
func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T, Value: %v\n", needInt(Small), needInt(Small))
	fmt.Printf("Type: %T, Value: %v\n", needFloat(Small), needFloat(Small))
}
