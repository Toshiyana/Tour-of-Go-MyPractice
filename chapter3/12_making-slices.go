// make関数によるスライスの作成：動的サイズの配列の作成が可能
// make関数：ゼロ化された配列を割り当て、その配列を指す指すスライスを返す（スライスは配列のポインタのようなもの）
// makeの第二引数：スライスのlength
// makeの第三引数：スライスのcapacity
// a := make([]int, 5) // len(a)=5
// b := make([]int, 0, 5) // len(b)=0, cap(b)=5
// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:]      // len(b)=4, cap(b)=4

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
