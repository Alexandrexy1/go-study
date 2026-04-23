package main

import (
	"fmt"
)

func main() {
	// fmt.Println(swap(3, 4)) // 0, 3

	// a, b := swap(40, 20)

	// fmt.Println(a, b) // 2, 0

	// result := plus(3)(1)

	// fmt.Println(result) // 4

	// num1, num2, num3 := plusAll(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// fmt.Println(num1, num2, num3) // 46, 47, 48

	// num := 10234
	// fmt.Println(strconv.FormatInt(int64(num), 10)) // 10234

	// const num1 = 10 // untyped constant, ela assume o tipo exigido pelo contexto (ex: int, int32, int64 ou até float)

	arr := [5]int{1, 2, 4: 500}
	arrStr := [10]string{5: "Olá, você quem está acompanhando os códigos :D"}

	fmt.Println(arr)    // [1, 2, 0, 0, 500]
	fmt.Println(arrStr) // [     Olá, você quem está acompanhando os códigos :D    ]

}

// func swap(a, b int) (res, rem int) {
// 	res = a / b
// 	rem = a % b
// 	return res, rem
// }

// func plus(a int) func(b int) int {
// 	return func(b int) int {
// 		return a + b
// 	}
// }

// func plusAll(nums ...int) (num1, num2, num3 int) {
// 	var out int
// 	for _, n := range nums {
// 		out += n
// 	}
// 	num1 = out + 1
// 	num2 = out + 2
// 	num3 = out + 3

// 	return num1, num2, num3
// }
