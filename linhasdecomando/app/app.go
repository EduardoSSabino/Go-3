package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Após a criação do programa, os comandos do terminal terão flag!
//go fun main.go ip --host amazon.com.br

//go fun main.go servidores --host devbook.com.br

// Gerar vai retorna a aplicação de linha de comadno pronta para ser executada
func Gerar() *cli.App { //é uma boa ideia ver a documentação co cli!
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na intenet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	// net, vai buscar os IPs

	ips, erro := net.LookupIP(host) //vai nos retornar uma slice ips e um erro
	if erro != nil {                //caso tenha um erro, aqui irei tratá-lo
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //uma função do pacote net que irá nos retornar os servidores. NS = Name Server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host) //o .Host é pra tirar do formato de struct
	}
}

/*Agora que eu ja tenho dois comandos executáveis, irei dar um go build no terminal
para poder compilar tudo isso que a gente fez, em um arquivo executavel
para podermos fazer a execução por esse arquivo*/

/*Após isso, aciono o comando ./linhadecomadno ip --host amazon.com.br
pq agora o arquivo ja esta compilado! e poderia fazer o mesmo para o metodo de descobrir os servidores*/
