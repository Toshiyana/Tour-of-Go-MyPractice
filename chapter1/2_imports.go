package main

import (
	"fmt"
	"math"
)

func main() {
	// Println: 改行あり、Printf：改行なし、値を途中で入れられる%
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// %gには入らない
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
}
