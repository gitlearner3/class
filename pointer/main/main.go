package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero(x)
	fmt.Println(x) // x is still 5
	zerop(&x)
	fmt.Println(x) // x will be 0 now

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}

func zero(x int) {
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	x = 0
}

func zerop(x *int) {
	fmt.Printf("%p\n", x)
	fmt.Println(*x)
	*x = 0
}
