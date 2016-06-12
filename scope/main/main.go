package main

import "fmt"

var x int = 42

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
	fmt.Println(z)
	z := 19
}

func foo() {
	fmt.Println(x)
	fmt.Println(y)
}
