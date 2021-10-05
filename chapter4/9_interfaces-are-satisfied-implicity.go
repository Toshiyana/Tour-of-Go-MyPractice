// インターフェースの暗黙的実行
// あんまり理解できていない

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// type T が interface I を実行するメソッド
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
