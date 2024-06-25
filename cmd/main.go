package main

import "fmt"

func soma(a, b float64) float64 {
	return a + b;
}

func subtracao(a, b float64) float64 {
	return a - b;
}

func multiplicacao(a, b float64) float64 {
	return a * b;
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Não é possível dividir por zero")
	}
	return a / b, nil
}

func operacao (a, b float64, op string) (float64, error) {
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
	for {
		fmt.Println("Digite a operação desejada (+)(-)(*)(/) ou 'sair' para encerrar a calculadora: ")
		fmt.Scanln(&op)
		if op == "sair" {
			fmt.Println("Saindo da calculadora")
			break
		} else {
		fmt.Println("Digite o primeiro número: ")
		fmt.Scanln(&a)
		fmt.Println("Digite o segundo número: ")
		fmt.Scanln(&b)
		resultado, err := operacao(a, b, op)
		if err != nil {
			fmt.Println("Erro", err)
		} else {
			fmt.Println(a, op, b, "=", resultado)
			}
		}
	}
}