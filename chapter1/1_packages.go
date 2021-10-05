// goプログラムはpackageで構成
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 擬似乱数の生成
	fmt.Println("My favorite number is", rand.Intn(10))
	// 実行毎に生成する乱数を変えるには、seed値を変える必要あり
	fmt.Println("Seed value is", rand.Seed)
}
