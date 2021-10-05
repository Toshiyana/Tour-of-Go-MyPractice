// map: キーと値をもつ。（pythonの辞書型みたいなもの）
// mapのゼロ値はnil.
// make()で型を指定して初期化が可能。(var部分でも初期化は可能)

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
