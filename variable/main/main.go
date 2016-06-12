package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v with type %T \n", a, a)
	fmt.Printf("%v with type %T \n", b, b)
	fmt.Printf("%v with type %T \n", c, c)
	fmt.Printf("%v with type %T \n", d, d)

	fmt.Printf("%v with type %T \n", e, e)
	fmt.Printf("%v with type %T \n", f, f)
	fmt.Printf("%v with type %T \n", g, g)
	fmt.Printf("%v with type %T \n", h, h)
}
