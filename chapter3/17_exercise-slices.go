// Pic関数の実装
// return: [][]uint8のgrayscale画像
// parameters
// dy: 画像の列数
// dx: 画像の行数
// 長さdyのスライスの各要素に、長さdxのスライスを割り当てることで画像を生成

package main

import "golang.org/x/tour/pic"
// picのソースコード(https://pkg.go.dev/golang.org/x/tour/pic)

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		img[y] := make([]uint8, dx)
		for x:= 0; x < dx; x++ {
			img[y][x] = uint8((x+y)/2)
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
	//pic.Show(Pic(256, 256)) //エラー。picのソースコードを見るとわかるが、dx, dyはconstで256が与えられている。
}