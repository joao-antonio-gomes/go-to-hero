package main

import (
	"fmt"
)

func main() {
	versao := 1.1
	fmt.Println("Bem vindo")
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", &comando)

	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Comando não reconhecido...")
	}
}
