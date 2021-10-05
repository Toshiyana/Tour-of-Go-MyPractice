// スライス＝配列への参照（ポインタみたいなもの）
// スライスは、元の配列の部分列を指しているので（ポインタ）、
// スライスの要素を変更すると、元の配列の要素も変わる。

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
