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

func divisao(a, b flaot64) float64 {
	if b != 0 {
		return a / b;
	} else {
		return fmt.Println("Não é possível dividir por zero!")
	}
}

func operacao (a, b flaot64, op string) float64 {
	switch op {
	case "+":
		return soma(a, b)
	case "-":
		return subtracao(a, b)
	case "*":
		return multiplicacao(a, b)
	case "/":
		return divisao(a, b)
	}
}

func main() {

	var a, b float64
	fmt.Println("Digite a operação desejada (+)(-)(*)(/): ")
	operacao(a, b, op)
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo número: ")
	fmt.Scanln(&b)
	fmt.Println("Resultado:", resultado)
}