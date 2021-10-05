// スライス

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// スライスは[]で定義。
	var s []int = primes[1:4] //1 ~ 3要素(pythonと同じ)
	fmt.Println(s)
}
