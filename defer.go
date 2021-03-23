package main

import "fmt"

func primero() {
	fmt.Println("primero")
}

func segundo() {
	fmt.Println("segundo")
}

func main() {
	defer primero() // se ejecuta antes de terminar la función
	segundo()
}
