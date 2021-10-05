package main

import "fmt"

func main() {
	fmt.Println("counting")

	//deferで渡した関数はstackされる。LIFO(last-in-fast-out)の順で実行される。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
