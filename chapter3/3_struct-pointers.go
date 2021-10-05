// structのfieldは、structのポインタを通してもアクセス可能。

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v   // ポインタを代入
	p.X = 1e9 //ポインタを通して,field Xの値を変更（*p).Xでもアクセス可能。
	fmt.Println(v)
	fmt.Println(p.X)
	fmt.Println((*p).X)
}
