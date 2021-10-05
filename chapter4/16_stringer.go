// Stringers
// Stringerは、fmtパッケージに定義されている最もよく使われているinterfaceの一つ

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age) //Sprintfであることに注意
}

func main() {
	a := Person{"Kenta", 14}
	b := Person{"Yuji", 25}
	fmt.Println(a, b)
}
