// Daniel R. Alvarenga - 09 de novembro de dois mil e vinte e três

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Chame aqui o exercício que quiser para ver sua execução

	ex_16()

}

func ex_0() {
	fmt.Println("Hello, World!")
}

func ex_1() {

	var n1 float64
	var n2 float64

	fmt.Print("Digite o primeiro número: ")
	if _, err := fmt.Scan(&n1); err != nil {
		fmt.Println("Erro ao receber o 1º número, verifique sua entrada: ", err)
		return
	}

	fmt.Print("Digite o segundo número: ")
	if _, err := fmt.Scan(&n2); err != nil {
		fmt.Println("Erro ao receber o 2º número, verifique sua entrada: ", err)
		return
	}

	sum := n1 + n2

	fmt.Println("A soma dos número inseridos é: ", sum)
}

func ex_2() {

	var n1 float64
	var n2 float64
	var n3 float64

	fmt.Print("Digite o primeiro número: ")
	if _, err := fmt.Scan(&n1); err != nil {
		fmt.Println("Erro ao receber o 1º número, verifique sua entrada: ", err)
		return
	}

	fmt.Print("Digite o segundo número: ")
	if _, err := fmt.Scan(&n2); err != nil {
		fmt.Println("Erro ao receber o 2º número, verifique sua entrada: ", err)
		return
	}

	fmt.Print("Digite o terceiro número: ")
	if _, err := fmt.Scan(&n3); err != nil {
		fmt.Println("Erro ao receber o 3º número, verifique sua entrada: ", err)
		return
	}

	media := (n1 + n2 + n3) / 3

	fmt.Printf("\nA média dos números inseridos é: %.2f", media)

}

func ex_3() {

	var n1 int

	fmt.Print("Insira um número inteiro: ")
	if _, err := fmt.Scan(&n1); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
		return
	}

	mol := n1 % 2

	if mol != 0 {
		fmt.Println("O número é ímpar")
	} else {
		fmt.Println("O número é par")
	}

}

func ex_4() {

	var n1 float64
	var n2 float64
	var op int

	fmt.Print("\nDigite o primeiro número: ")
	if _, err := fmt.Scan(&n1); err != nil {
		fmt.Println("Erro ao receber o 1º número, verifique sua entrada: ", err)
		return
	}

	fmt.Print("Digite o segundo número: ")
	if _, err := fmt.Scan(&n2); err != nil {
		fmt.Println("Erro ao receber o 2º número, verifique sua entrada: ", err)
		return
	}

	fmt.Print("\n 1 - Soma")
	fmt.Print("\n 2 - Subtração")
	fmt.Print("\n 3 - Multiplicação")
	fmt.Print("\n 4 - Divisão")
	fmt.Print("\n\nDigite a opção deseja: ")
	if _, err := fmt.Scan(&op); err != nil {
		fmt.Println("Erro ao receber o 1º número, verifique sua entrada: ", err)
		return
	}

	if op > 0 && op < 5 {

		if op == 1 {
			valor := n1 + n2
			fmt.Printf("%.2f + %.2f é gual a %.2f", n1, n2, valor)
		} else if op == 2 {
			valor := n1 - n2
			fmt.Printf("%.2f - %.2f é gual a %.2f", n1, n2, valor)
		} else if op == 1 {
			valor := n1 * n2
			fmt.Printf("%.2f * %.2f é gual a %.2f", n1, n2, valor)
		} else {
			valor := n1 / n2
			fmt.Printf("%.2f / %.2f é gual a %.2f", n1, n2, valor)
		}

	} else {
		fmt.Println("Você não inseriu uma opção válida, tente novamente...")
	}
}

func ex_5() {
	var n int

	fmt.Print("\nInsira um número inteiro: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
	}

	sum := 0

	for i := n; i > 0; i-- {
		sum += i
	}

	fmt.Printf("\n%d! é %d", n, sum)
}

func ex_6() {
	var n int

	fmt.Print("\nInsira um número inteiro: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
	}

	primo := true

	if n < 2 {
		fmt.Printf("\n%d é um número primo.", n)

		return
	}

	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			primo = false
		}
	}

	if primo {
		fmt.Printf("\n%d é um número primo.", n)
	} else {
		fmt.Printf("\n%d não é um número primo.", n)
	}

}

func ex_7() {
	var n float64

	fmt.Print("\nInsira um valor em graus Celsius: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o valor, verifique sua entrada: ", err)
	}

	f := ((9.0 / 5.0) * n) + 32

	fmt.Printf("\n%.2f °C são equivalentes a %.2f °F", n, f)
}

func ex_8() {
	var n int

	fmt.Print("\nInsira um número inteiro: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
	}

	sum := 0

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			sum += i
		}
	}

	fmt.Printf("\nA soma dos números ímpares de 1 até %d é %d", n, sum)
}

func ex_9() {
	var n float64

	fmt.Print("\nInsira um número: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
	}

	fmt.Printf("\nTabuada do %2.f\n", n)

	for i := 1; i <= 10; i++ {
		p := n * float64(i)
		fmt.Printf("\n%.2f × %d = %2.f", n, i, p)
	}

}

func ex_10() {
	var n int

	fmt.Print("\nInsira um número inteiro: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
	}

	n1 := 0
	n2 := 1

	fmt.Printf("Primeiros %d números da sequência de Fibonacci:\n", n)

	for i := 0; i <= n; i++ {
		d := n1 + n2
		fmt.Printf("%d\n", d)

		temp := n2
		n2 = n1 + n2
		n1 = temp
	}

}

func ex_11() {
	var palavra string

	fmt.Print("\nEscreva uma palavra: ")
	fmt.Scan(&palavra)

	runas := []rune(palavra)
	runas2 := []rune(palavra)
	len := len(palavra)

	for i := 0; i < len; i++ {
		runas[i] = runas2[len-1-i]
	}

	fmt.Printf("A palavra %s invertida fica %s", palavra, string(runas))
}

func ex_12() {
	var palavra string

	fmt.Print("\nEscreva uma palavra: ")
	fmt.Scan(&palavra)

	runas := []rune(palavra)
	runas2 := []rune(palavra)
	len := len(palavra)

	for i := 0; i < len; i++ {
		runas[i] = runas2[len-1-i]
	}

	palavra_invertida := string(runas)

	if palavra == palavra_invertida {
		fmt.Printf("A palavra %s é um palíndromo", palavra)
	} else {
		fmt.Printf("A palavra %s não é um palíndromo", palavra)
	}

}

func ex_13() {
	var n int

	fmt.Print("\nInsira um número inteiro: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
	}

	n_s := string([]rune(strconv.Itoa(n)))
	len := len(n_s)

	fmt.Printf("O número %d possui %d dígito(s)", n, len)

}

func ex_14() {
	var n int

	fmt.Print("\nInsira um número inteiro: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
		return
	}

	vetor := []int{1}

	for i := 1; i <= n; i++ {

		for a := 0; a < len(vetor); a++ {
			fmt.Printf("%d ", vetor[a])
		}

		fmt.Printf("\n")

		temp := make([]int, len(vetor)+1)

		temp[0], temp[len(temp)-1] = 1, 1

		for b := 1; b < len(vetor); b++ {
			temp[b] = vetor[b-1] + vetor[b]
		}

		vetor = temp
	}

}

func ex_15() {
	var n float64

	fmt.Print("\nInsira um valor em reais: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
		return
	}

	cotacao := 4.94

	fmt.Printf("R$%.2f equivalem a $%.2f, levando em consideração uma cotação de R$%.2f", n, n*cotacao, cotacao)
}

func ex_16() {
	fmt.Println("\nIremos montar uma lista, quando quiser finalizar insira -1")

	list := make([]float64, 0)

	n := 1.0

	for n != -1 {
		fmt.Print("Insira um número: ")
		if _, err := fmt.Scan(&n); err != nil {
			fmt.Println("Erro ao receber o número, verifique sua entrada: ", err)
			return
		}

		list = append(list, n)
	}

	ordenada := false

	for ordenada == false {
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}

		ordenada = true

		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				ordenada = false
			}
		}
	}

	fmt.Print("\nLista ordenada: ", list)
}
