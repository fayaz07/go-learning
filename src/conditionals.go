package main

import "fmt"

func main() {
	x := 10
	y := 14

	if x < y {
		fmt.Println(x)
	} else if x == y {
		fmt.Println("equal")
	} else {
		fmt.Println(y)
	}

	switch x {
	case 10:
		fmt.Println("case 1")
		break
	case 14:
		fmt.Println("case 2")
		break
	}
}
