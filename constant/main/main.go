package main

import "fmt"

const p string = "Constant #1"
const q = 42

func main() {
	const x = "Constant local #1"
	const y int = 43

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("x - ", x)
	fmt.Println("y - ", y)
}
