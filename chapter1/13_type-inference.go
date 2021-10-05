package main

import "fmt"

func main() {
	// := の暗黙的宣言を行う場合、右側の変数から型推論するので、右側には値が入っている必要あり
	var i int = 5
	j := i

	fmt.Println(i, j)
}
