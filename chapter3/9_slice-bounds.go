// スライスするときの上限、下限の省略（pythonと同じ）
// 以下の配列と
// var a [10]int
// 以下のスライス式は等価
// a[0:10]
// a[:10]
// a[0:]
// a[:]

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
