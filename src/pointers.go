package main

import "fmt"

func main() {
	a := 5
	b := &a

	println(a, b)
	fmt.Printf("%T\n", b)

	// values
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with pointer
	*b = 10
	fmt.Println(a)
}
