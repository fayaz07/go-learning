package main

import "fmt"

func main() {
	// var arrayEx [10]int32

	// arrayEx[0] = 1
	// arrayEx[1] = 2

	// arrayEx = [10]int32{1, 2, 3, 4}

	// fmt.Println(arrayEx)

	var sliceEx []int32
	fmt.Println(sliceEx)

	sliceEx = []int32{1, 2, 3}
	fmt.Println(sliceEx, sliceEx[2])
	fmt.Println(len(sliceEx))
	fmt.Println(sliceEx[0:2])
}
