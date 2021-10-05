// Type switches
// 通常のswitchと異なり、型switchのcaseは型を指定する。
// 指定されたインターフェースの値が保持する値の型と比較される、

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) { // キーワードtypeを指定
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
