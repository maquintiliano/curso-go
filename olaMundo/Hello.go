package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitoramentos = 3
	delay          = 5
)

func main() {

	exibeIntroducao()

	for {

		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {

	//declarando variaveis
	nome := "Maria"
	versao := 1.1

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)

	//descobrir o tipo da variavel usando reflect
	//fmt.Println("O tipo da variavel versão é", reflect.TypeOf(versao))

}

func lerComando() int {

	//capturando o input do usuario e armazenando em uma variavel
	//& indica o endereço da variavel comando. Funciona como um ponteiro
	var lerComando int
	fmt.Scan(&lerComando)

	fmt.Println("O comando escolhido foi", lerComando)

	return lerComando
}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir log")
	fmt.Println("0- Sair do programa")
	fmt.Println("")

}

func iniciarMonitoramento() {

	fmt.Println("Monitorando...")
	fmt.Println("")

	//slice
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		//usando range para retornar a posição e o valor da posição
		for i, sites := range sites {
			fmt.Println("Testando site", i, ":", sites)
			testaSite(sites)
			fmt.Println("")
		}
	}
	time.Sleep(delay * time.Second)
	fmt.Println("")
}

func testaSite(site string) {
	//acessando o site
	//o comando get retorna dois retornos, porém caso eu queira
	//para lidar com apenas um deles, podemos utilizar o undeline no lugar do outro
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code", resp.StatusCode)
		registraLog(site, false)

	}

}

func leSitesDoArquivo() []string {

	var sites []string

	//abrir o arquivo de sites.txt e detectar caso algum erro aconteça na abertura
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		//função para ler cada linha ReadString()
		linha, err := leitor.ReadString('\n')
		//função trimSpace tira os espaços, tabs e /n do fim da linha
		linha = strings.TrimSpace(linha)
		//função append insere a linha dentro do slice sites
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	//fechar o arquivo
	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {

	//cria, abre, lê e escreve um novo arquivo
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {

		fmt.Println(err)

	}
	//utilizando strconv para converter bool em string
	//format time tem formatação diferente das demais (procurar no google)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "-online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	//abrir o arquivo de log.txt e detectar caso algum erro aconteça na abertura
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Exibindo logs...")

	fmt.Println(string(arquivo))

}
