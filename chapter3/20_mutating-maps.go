// mapの操作

// m へ要素(elem)の挿入や更新:
// m[key] = elem

// 要素の取得:
// elem = m[key]

// 要素の削除:
// delete(m, key)

// キーに対する要素が存在するかどうかは、2つの目の値で確認
// elem, ok = m[key]
// キーが存在する時、okはtrue. 存在しない時、okはfalseで、値はゼロ値。

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
