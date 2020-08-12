package main

import "fmt"

func main() {
	fmt.Println(add(10, 100))
	fmt.Println(max(10, 100))
	printSomething()

	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}

func add(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func printSomething() {
	fmt.Println("I printed something")
}

func swap(x, y string) (string, string) {
	return y, x
}
