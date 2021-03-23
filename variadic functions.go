import "fmt"

func sumar(args ...int) int { // args es un slice de los parámetros recibidos
	total := 0
	for _, v := range args { // iterar sobre los parámetros como un slice
		total += v
	}

	return total
}

func main() {
	a := []int{1, 4, 2}
	fmt.Println(sumar(4, 5, 2))
	fmt.Println(sumar(a...)) // cada elemento del slice se envía como parámetros
}
