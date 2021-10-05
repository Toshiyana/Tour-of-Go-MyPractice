// Readers
// ioパッケージ：データストリームを読む時に利用

// 与えられたデータのバイトサイズとエラーの値を返す

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	// 8byte毎に読み出し
	b := make([]byte, 8)
	for { //無限ループ
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
