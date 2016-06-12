package main

import "fmt"

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
	fmt.Println(increment())
	fmt.Println(increment())

	z := 0
	increment2 := func() int {
		z++
		return z
	}
	fmt.Println(increment2())
	fmt.Println(increment2())
}

func foo() {
	fmt.Println(x)
}

func increment() int {
	x++
	return x
}

var x int = 42
