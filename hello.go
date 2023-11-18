package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
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
		for siteNumber, siteName := range getSitesFromFile() {
			verifySiteIsOnline(siteName, siteNumber)
		}
		time.Sleep(monitoringDelay)
		fmt.Println("")
	}
}

func verifySiteIsOnline(siteName string, number int) {
	site := getSiteToDoRequest()
	resp, err := http.Get(site)
	handleErr(err)

	fmt.Println("Testando site", number, ":", siteName)
	if resp.StatusCode == 200 {
		fmt.Println("Site", siteName, "foi carregado com sucesso.")
	} else {
		fmt.Println("Site", siteName, "está com problemas. Status code:", resp.StatusCode)
	}
}

func getSitesFromFile() []string {
	//assim lemos o arquivo inteiro
	//file, err := os.ReadFile("sites")

	//assim apenas abrimos o arquivo
	file, err := os.Open("sites")
	handleErr(err)

	reader := bufio.NewReader(file)

	var sites []string
	for {
		line, err := reader.ReadString('\n')

		if handleErr(err) && err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	file.Close()

	return sites
}

func handleErr(err error) bool {
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return true
	}
	return false
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
