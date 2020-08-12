package main

func main() {
	arr := []int{12, 15, 16, 80, 32, 43}

	for i, el := range arr {
		println(i, el)
	}

	for _, el := range arr {
		println(el)
	}

	emails := map[string]string{"john": "john@doe.com"}

	for k, v := range emails {
		println(k, v)
	}
}
