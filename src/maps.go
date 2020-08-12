package main

import "fmt"

func main() {
	// define map
	email := make(map[string]string)

	// assign key values
	email["john"] = "john@doe.com"
	email["jane"] = "jane@doe.com"

	fmt.Println(email)
	fmt.Println(len(email))

	// delete
	delete(email, "jane")

	fmt.Println(email)
	fmt.Println(len(email))

	// declare and initialize
	emails := map[string]string{"john": "john@doe.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
}
