// 定数宣言
// 定数は文字(character)、文字列(string)、boolean、数値(numeric)のみで使用可能。
// 定数宣言では、:= は使えない。

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
