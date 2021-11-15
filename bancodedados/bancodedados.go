package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//
/*No terminal foi usado:
-go mod init banco-de-dados
-go get github.com/go-sql-driver/mysql*/

func main() {
	stringConexao := "eduardo:@Sabiedusilvawheel13@/bancodedaos?charset=utf8&parseTime=True&loc=local" //"usuario:senha@/banco"
	db, erro := sql.Open("mysql", stringConexao)                                                       //abre o banco de dados
	if erro != nil {                                                                                   //caso tenha um erro
		log.Fatal(erro) //para tudo por aqui e imprimi uma mensagem dizendo qual foi o erro
	}
	defer db.Close() //fecho o meu banco de dados de qualquer maneira

	if erro = db.Ping(); erro != nil { //.Ping faz um teste pra ver se o programa está conectado com o banco de dados
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	//retronar dados pra gente

	linhas, erro := db.Query("selecy * from usuarios") // query : É uma solicitação de informações feita ao banco de dados. "Selecionar a partir de usuarios".
	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()
	fmt.Println(linhas)
}
