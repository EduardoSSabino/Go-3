package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template //contem todos os templates da nossa aplicação
type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html")) //vai pegar todos meus arquivos html e vai jogar dentro do meu template

	u := usuario{
		"João",
		"joao.pedro@gmail.com",
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})
	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
