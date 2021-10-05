// switch文
// if-elseより短く書ける
// Goのswitchの特徴
// ・他の言語とは異なり、breakが自動的に提供される（書く必要ない）
// ・caseは定数、整数である必要ない。

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
