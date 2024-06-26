package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func soma(a, b float64) float64 {
	return a + b
}

func subtracao(a, b float64) float64 {
	return a - b
}

func multiplicacao(a, b float64) float64 {
	return a * b
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Não é possível dividir por zero")
	}
	return a / b, nil
}

func operacao(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return soma(a, b), nil
	case "-":
		return subtracao(a, b), nil
	case "*":
		return multiplicacao(a, b), nil
	case "/":
		resultado, err := divisao(a, b)
		if err != nil {
			return 0, err
		}
		return resultado, nil
	case "sair":
		return 0, fmt.Errorf("Saindo do programa")
	default:
		return 0, fmt.Errorf("Operação inválida")
	}
}

func main() {

	var a, b float64
	var op string
	var err error
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Digite a operação desejada (+)(-)(*)(/) ou 's' para encerrar a calculadora: ")
		op, _ = reader.ReadString('\n')
		op = strings.TrimSpace(op)
		if op != "+" && op != "-" && op != "*" && op != "/" && op != "sair" {
			fmt.Println("Operação inválida.")
			continue

		} else if op == "s" {
			fmt.Println("Saindo da calculadora")
			break
		} else {
			fmt.Println("Digite o primeiro número: ")
			_, err = fmt.Scanf("%f", &a)
			if err != nil {
				fmt.Println("Entrada inválida. Tente novamente.")
				reader.ReadString('\n') // limpar o buffer de entrada
				continue
			}
			fmt.Println("Digite o segundo número: ")
			_, err = fmt.Scanf("%f", &b)
			if err != nil {
				fmt.Println("Entrada inválida. Tente novamente.")
				reader.ReadString('\n') // limpar o buffer de entrada
				continue
			}
			resultado, err := operacao(a, b, op)
			if err != nil {
				fmt.Println("Erro", err)
			} else {
				fmt.Println(a, op, b, "=", resultado)
			}
		}
	}
}
