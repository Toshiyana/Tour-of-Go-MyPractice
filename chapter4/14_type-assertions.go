// インターフェース型が特定の型を保持しているかは、
// t, ok := i.(T)
// として、iがTを保持していれば、tは基になる値、okはtrueになる。

package main

import "fmt"

func main() {
	var i interface{} = "hello"
	// var i interface{} = 0.5

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
