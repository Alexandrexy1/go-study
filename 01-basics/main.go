package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// var i int
	// for i < 10 {
	// 	fmt.Println("Infinito") // como não aumento o valor da var i, o Println roda infinitamente
	// }

	// arr := [10]int{1, 4, 10, 9, 81, 6, 26, 28, 37, 75}
	// for _, el := range arr {
	// 	fmt.Println(&el, el) // mostra o endereço de memória e o elemento
	// }

	if x := math.Sqrt(81); x < 10 { // short statement é o termo da var dentro do if, só existe naquele escopo do if
		fmt.Println("passou")
	}

	do(1) // retorno: um dois

	isWekeend(time.Now()) // não é fim de semana

	switch x := math.Sqrt(16); x {
	case 4:
		fmt.Println("O resultado é 4")
	default:
		fmt.Println("Resultado diferente de 4")
	}
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println("um")
		fallthrough // ele passa para o próximo case
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	default:
		fmt.Println("valor padrão")
	}
}

func isWekeend(x time.Time) {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		fmt.Println("Não é fim de semana")
	default:
		fmt.Println("É fim de semana")
	}
}
