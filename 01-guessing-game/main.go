package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello! Let's play jogo da advinhação")
	fmt.Println("Advinhe o número sorteado. Você terá 10 chances, o número inteiro, aleatório e vai de 0 a 100")

	scanner := bufio.NewScanner(os.Stdin)
	chances := [10]int64{}

	num := rand.Int64N(101)

	for i := range chances {
		fmt.Println("Qual o seu chute? ")
		scanner.Scan()
		chance := scanner.Text()

		chance = strings.TrimSpace(chance)

		chanceInt, err := strconv.ParseInt(chance, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro")
		}

		switch {
		case chanceInt > num:
			fmt.Println("Você errou. Número sorteado é menor que o número ", chanceInt)
		case chanceInt < num:
			fmt.Println("Você errou. Número sorteado é maior que o número ", chanceInt)
		case chanceInt == num:
			fmt.Println("Parabéns! Você acertou!!")
			fmt.Println("essa foi todas as suas tentativas até acertar: ", chances[:i]) // 50, 25, 30 foi a sequência utilizada (30 era o número sorteado)
			return
		}

		chances[i] = chanceInt
	}

	fmt.Println("essa foi todas as suas tentativas: ", chances)
}
