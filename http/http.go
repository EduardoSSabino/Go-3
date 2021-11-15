package main

import (
	"log"
	"net/http"
)

// HTTP é um protocolo de comunicação,  base da comunicação web
//o http funciona da seguinte forma, o cliente faz uma requisição ao servidor esse servidor processa a requisição e devolve uma resposta

/*exemplo, imagine que você entre num site e você quer fazer uma conta nele. Você entra no site preenche seus dados e faz uma
requisição pro servidor, no caso o site, que vai processar os dados que você inseriu e te dará uma resposta abseado nesse processamento
então ele pode te dar uma resposta faladno que seu cadastro foi concluído com sucesso ou que teve algum erro no seu dado, ou seja,
sempre teremos esse fluxo, você faz uma requisição ao servidor, o servidor processa e devolve uma resposta.*/

//request - requisição feita pelo cliente

//response -  resposta do servidor

/*rotas - são uma maneira de você conseguir identificar um tipo de mensagem que ta sendo enviado e apartir disso identificar que tipo de processamento
o servidor vai te que fazer em cima dessa mensagem.
			URI - identificador do recurso. Uma forma de falar pro servidor o que estou falando,  "isso aqui tem haver com produto"
			Método - GET, POST, PUT, DELETE
			         GET : pegar dados de um usuario
					 POST : normalmente é usado para cadastrar dados de um usuario novo
					 PUT : usuado para atualizações de dados
					 DELETE : usado para deletar dados*/

func home(w http.ResponseWriter, r *http.Request) { /*esse método recebe dois parametro, primeiro é justamente o URI da rota
	e o segundo um função que recebe a requisição e sabe lidar com ela*/
	w.Write([]byte("Olá mundo!")) //nesse caso não podemos usar fmt.Println, pra imprimir algo, faremos desse jeito que fizemos.
}

func main() {

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil)) /*passamos uma porta (:5000) pro sistema que vai ficar aberta, e é por ela que vamos ficar ouvindo
	as requisições e dando as respostas*/
}

/* A REQUISIÇÃO NÓS FAREMOS PELO O NAVEGADOR! USANDO localhost:5000 + URI*/
