// Go言語の型は以下。
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 の別名

// rune // int32 の別名
//      // Unicode のコードポイントを表す

// float32 float64

// complex64 complex128

// int, uint, uintptr 型は、32-bitのシステムでは32 bitで、64-bitのシステムでは64 bit.
// サイズ、符号なし( unsigned )整数の型を使うための特別な理由がない限り、整数の変数が必要な場合は int を使う

package main

import (
	"fmt"
	"math/cmplx"
)

// varもimportと同様にまとめて宣言可能
var (
	ToBe   bool       = false
	MaxInt uint64     = 1 << 3               //ビット演算子、左シフト<<
	z      complex128 = cmplx.Sqrt(-5 + 12i) //複素数
)

func main() {
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}
