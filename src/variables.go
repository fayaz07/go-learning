package main

import "fmt"

var name string
var a int
var i32 int32
var i64 int64
var checkbool bool
var f32 float32
var f64 float64

func main() {
	fmt.Println("Variables before initializing")
	printVariables()

	name = "fayaz"
	a = 5
	i32 = 999999999
	i64 = 999999999999999999

	f32 = 99999999999999999999999999999999999999.999
	f64 = 99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.999

	checkbool = true
	fmt.Println("\nVariables after initializing")
	printVariables()
}

func printVariables() {
	fmt.Printf("Name: %v\n", name)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("i32: %v\n", i32)
	fmt.Printf("i64: %v\n", i64)
	fmt.Printf("checkbool: %v\n", checkbool)
	fmt.Printf("f32: %v\n", f32)
	fmt.Printf("f64: %v\n", f64)
  fmt.Print(f64);
}
