package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// var i int
	// for i < 10 {
	// 	fmt.Println("Infinito") // como não aumento o valor da var i, o Println roda infinitamente
	// }

	arr := [10]int{1, 4, 10, 9, 81, 6, 26, 28, 37, 75}
	for _, el := range arr {
		fmt.Println(&el, el) // mostra o endereço de memória e o elemento
	}
}
