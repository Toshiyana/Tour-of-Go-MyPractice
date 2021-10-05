// strct(構造体)：fieldの集まり。
// structのfieldはドット.でアクセス

package main

import "fmt"

type Vertex struct {
	// X, Yはstrct Vertexのfield.
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	// ドットでfieldにアクセス
	v.X = 4
	fmt.Println(v.X)

}
