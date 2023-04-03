package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas. Informe ao usuário o resultado.

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Qual é a string 1 :")
	scanner.Scan()
	frase1 := scanner.Text()

	frase1 = strings.ToUpper(frase1)

	fmt.Println("A frase em maiusculo fica", frase1)

}
