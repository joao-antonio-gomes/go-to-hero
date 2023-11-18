package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const monitoringDelay = 5 * time.Second
const monitoringQuantity = 5

func main() {
	showIntro()

	for {
		showMenu()

		command := readCommand()
		fmt.Println("")

		switch command {
		case 1:
			initializeMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido...")
			os.Exit(-1)
		}
	}
}

func initializeMonitoring() {
	fmt.Println("Monitorando...")

	for i := 0; i < monitoringQuantity; i++ {
		for siteNumber, siteName := range getSitesSlice() {
			verifySiteIsOnline(siteName, siteNumber)
		}
		time.Sleep(monitoringDelay)
		fmt.Println("")
	}
}

func verifySiteIsOnline(siteName string, number int) {
	site := getSiteToDoRequest()
	resp, _ := http.Get(site)

	fmt.Println("Testando site", number, ":", siteName)
	if resp.StatusCode == 200 {
		fmt.Println("Site", siteName, "foi carregado com sucesso.")
	} else {
		fmt.Println("Site", siteName, "está com problemas. Status code:", resp.StatusCode)
	}
}

func getSitesSlice() []string {
	return []string{"https://twitter.com", "https://instagram.com", "https://globo.com"}
}

func getSiteToDoRequest() string {
	status := "200"
	rand := rand.Int31n(2)
	if rand < 1 {
		status = "404"
	}

	site := "https://httpbin.org/status/" + status
	return site
}

func readCommand() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func showIntro() {
	versao := 1.1
	fmt.Println("Bem vindo")
	fmt.Println("Este programa está na versão", versao)
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}
