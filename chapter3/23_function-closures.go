// closure:
// それ自身の外部から変数を参照する関数値(関数を返す関数)

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int { //関数値をreturn
		sum += x
		return sum // sumはバインドされている
	}
}

func main() {
	pos, neg := adder(), adder() //関数値の代入
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
