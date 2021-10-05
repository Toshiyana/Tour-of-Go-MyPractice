// rangeの使い方
// _の利用：_をindexやvalueに用いた場合、捨てることができる (pythonとは異なり、_とすると捨てられるので出力できない)
// for i, _ := range pow
// for _, value := range pow
// valueの省略：indexだけが必要な場合、2つ目の値を省略
// for i := range pow

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		//fmt.Printf("%d\n", i) // index
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
