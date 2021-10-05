package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 書き方1
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// 書き方2（mapに渡すトップレベルの型が単純な型名のとき、リテラルから推論できるため省略可能）
var n = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  //型名省略
	"Google":    {37.42202, -122.08408}, //型名省略
}

func main() {
	fmt.Println(m)
	fmt.Println(n)
}
