package main

import (
	"bufio"
	"fmt"
	"go-challengers/mathutils"
	"os"
	"strconv"
)

func main() {

}

func exercises() {
	// # 1. Pacotes
	fmt.Println("Resultado da Soma:", mathutils.Soma(10, 5))

	// # 2. Nomes Públicos e Privados
	fmt.Println("Resultado da Multiplicação:", mathutils.Multiply(5, 2))

	// # 3. Funções
	result := Calc(5, 2, "multiplicacao")
	fmt.Printf("\nA multiplicação é: %d\n", result)

	// # 4. Variáveis
	name := "João"
	idade := 30
	saldo := 250.75
	fmt.Println("\nInformações do usuário:")
	fmt.Printf("Nome: %s\n", name)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Saldo: %.2f\n", saldo)

	// # 5. Arrays
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("\nElementos do Array:")
	for _, elements := range arr {
		fmt.Printf("%d\n", elements)
	}

	// 6. Loops
	for i := range 10 {
		if i%2 == 0 {
			fmt.Printf("\nPar: %d", i)
		} else {
			fmt.Printf("\nImpar: %d", i)
		}
	}

	// 7. Condicionais (IFs)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Qual é a nota?")
	scanner.Scan()

	nota := scanner.Text()

	// Convertendo para float64 em vez de int64
	notaFloat, err := strconv.ParseFloat(nota, 64)
	if err != nil {
		fmt.Println("Erro ao converter a nota. Por favor, insira um número válido.")
		return
	}

	if notaFloat > 10 || notaFloat < 0 {
		fmt.Println("Digite uma nota válida, de 0 a 10")
		return
	}

	// Verificando intervalos de nota
	if notaFloat < 5 {
		fmt.Println("Reprovado")
	} else if notaFloat >= 5 && notaFloat <= 6.9 {
		fmt.Println("Recuperação")
	} else {
		fmt.Println("Aprovado")
	}

	fmt.Println("\nInicio do programa")

	defer fmt.Println("\nFim do programa")

	scanner2 := bufio.NewScanner(os.Stdin)

	fmt.Println("\nQual é seu nome?")
	scanner.Scan()

	name2 := scanner2.Text()

	// Imprimindo o nome
	fmt.Printf("\nSeu nome é: %s\n", name2)

	// 9. Ponteiros
	var number int

	pointer := &number

	*pointer = 10

	fmt.Printf("\nO numero é: %d", number)

	*pointer = 20

	fmt.Printf("\nO numero é: %d", number)
}

func Calc(a, b int, operation string) int {
	if sum := a + b; operation == "soma" {
		return sum

	}

	if sub := a - b; operation == "subtracao" {
		fmt.Printf("A subtração é: %d", sub)

		return sub
	}

	if multi := a * b; operation == "multiplicacao" {
		return multi
	}
	if div := a / b; operation == "divisao" {
		fmt.Printf("A divisao é: %d", div)

		return div
	}

	// Retorno padrão caso a operação seja inválida
	fmt.Println("Operação inválida.")
	return 0

}
