package main

import (
	"bufio" // para leitura bufferizada
	"fmt"
	"os" // para manipulação do sistema operacional
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

func lerEntrada(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)     // cria um novo leitor de entrada, que lê a partir do terminal (os.Stdin)
	fmt.Print(prompt)                       // exibe o prompt para o usuário
	entrada, err := reader.ReadString('\n') // lê a entrada do usuário até que ele pressione a tecla Enter
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(entrada), nil // remove qualquer espaço em branco (icluindo nova linha) do inicio ao fim da string
}

func tratamentoSaida(a, b, resultado float64, op string) string {
	if a == float64(int64(a)) && b == float64(int64(b)) && resultado == float64(int64(resultado)) { // a,b e resultado são inteiros
		return fmt.Sprintf("%d %s %d = %d", int64(a), op, int64(b), int64(resultado))
	} else if a == float64(int64(a)) && b == float64(int64(b)) { // a e b são inteiros
		return fmt.Sprintf("%d %s %d = %f", int64(a), op, int64(b), resultado)
	} else if a == float64(int64(a)) && resultado == float64(int64(resultado)) { // a e resultado são inteiros
		return fmt.Sprintf("%d %s %f = %d", int64(a), op, b, int64(resultado))
	} else if b == float64(int64(b)) && resultado == float64(int64(resultado)) { // b e resultado são inteiros
		return fmt.Sprintf("%f %s %d = %d", a, op, int64(b), int64(resultado))
	} else if a == float64(int64(a)) { // somente a é inteiro
		return fmt.Sprintf("%d %s %f = %f", int64(a), op, b, resultado)
	} else if b == float64(int64(b)) { // somente b é inteiro
		return fmt.Sprintf("%f %s %d = %f", a, op, int64(b), resultado)
	} else if resultado == float64(int64(resultado)) { // somente resultado é inteiro
		return fmt.Sprintf("%f %s %f = %d", a, op, b, int64(resultado))
	}

	return fmt.Sprintf("%f %s %f = %f", a, op, b, resultado)
}

func main() {

	var a, b float64
	var op string
	var err error
	var historico []string
	var resultadoStr string

	for {
		op, err = lerEntrada("Digite a operação (+, -, *, /) ou 'sair' para encerrar calculadora: ")
		if err != nil {
			fmt.Println("Erro na leitura da entrada", err)
		}
		if op != "+" && op != "-" && op != "*" && op != "/" && op != "s" {
			fmt.Println("Operação inválida.")
			continue
		}

		if op == "s" {
			fmt.Println("Saindo da calculadora")
			break
		}

		fmt.Println("Digite o primeiro número: ")
		_, err = fmt.Scanf("%f", &a)
		if err != nil {
			fmt.Println("Entrada inválida. Tente novamente.")
			continue
		}

		fmt.Println("Digite o segundo número: ")
		_, err = fmt.Scanf("%f", &b)
		if err != nil {
			fmt.Println("Entrada inválida. Tente novamente.")
			continue
		}
		resultado, err := operacao(a, b, op)
		if err != nil {
			fmt.Println("Erro", err)
		} else {
			resultadoStr = tratamentoSaida(a, b, resultado, op)
			historico = append(historico, resultadoStr)
			fmt.Println(resultadoStr)
		}
	}
	fmt.Println("Histórico de Operações:")
	for _, h := range historico {
		fmt.Println(h)
	}
}
