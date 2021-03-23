package main

import "fmt"

func main() {
	suma := func(x, y int) int {
		return x + y
	}

	fmt.Println(suma(8, 9))
}
