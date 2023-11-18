package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const monitoringDelay = 5 * time.Second
const monitoringQuantity = 5

func main() {
	showIntro()
	sites := getSitesFromFile()
	for {
		showMenu()

		command := readCommand()
		fmt.Println("")

		switch command {
		case 1:
			initializeMonitoring(sites)
		case 2:
			printLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido...")
			os.Exit(-1)
		}
	}
}

func printLogs() {
	file, err := os.ReadFile("log.txt")
	handleErr(err)

	fileContent := string(file)
	newFileContent := strings.ReplaceAll(fileContent, ",", " - ")

	fmt.Println(newFileContent)
}

func initializeMonitoring(sites []string) {
	fmt.Println("Monitorando...")

	for i := 0; i < monitoringQuantity; i++ {
		for siteNumber, siteName := range sites {
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
		registerLog(siteName, true)
	} else {
		fmt.Println("Site", siteName, "está com problemas. Status code:", resp.StatusCode)
		registerLog(siteName, false)
	}
}

func registerLog(siteName string, online bool) {
	file, err := os.OpenFile("log.txt", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, 0666)
	handleErr(err)

	logTime := time.Now().Format("02/01/2006 15:04:05")
	logPhrase := logTime + "," + siteName + "," + strconv.FormatBool(online) + "\n"
	file.WriteString(logPhrase)

	file.Close()
}

func getSitesFromFile() []string {
	file, err := os.Open("sites.txt")
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
