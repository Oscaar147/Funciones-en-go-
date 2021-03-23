package main

import "fmt"

func promedio(arreglo []int) float64 { // se crea una copia del slice pero no de los elementos
	total := 0.0
	for _, v := range arreglo {
		total += float64(v)
	}

	return total / float64(len(arreglo))
}

func main() {
	a := []int{4, 4, 7}
	fmt.Println(promedio(a)) // se
}
