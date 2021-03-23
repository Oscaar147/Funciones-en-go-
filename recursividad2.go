package mai

import "fmt"

func generadorPares() func() uint {
	i := uint(0) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint {
		var par = i
		i += 2
		return par
	}
}

func main() {
	nextPar := generadorPares()
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
}
