package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//aqui ficará as funções que meu servidor irá processar

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario essa função irá criar um usuario e insere um novo usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	//quando no postman eu uso localhost:5000, eu quero cadastrar um usuario no banco
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body) //ioutil é muito usado pra fazer leitura de entrada e saída de dados. meu readAll vai ler o meu corpo que está dentro do "r" da assinatura
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	//Agora vamos insere o usuario joão no banco de dados!
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	// PREPARE STATEMENT : criamos um comando de inserção que é muito utilizado pra evitar o ataque chamado "sql injection"

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)") // estamos tentando evitar um ataque!
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email) //Exec executa uma instrução preparada com os argumentos fornecidos e retorna um Resultado resumindo o efeito da instrução.
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserdo!"))
		return
	}

	// STATUS CODES : indica o que aconteceu com a requisição, se ela deu certo ou deu problema
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso Id: %d", idInserido)))
}

// BuscarUsuarios : traz todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
	}
	defer db.Close()

	// SELECT * FROM USUARIOS : "selecione tudo dos usuarios"

	//Pra consultar usamos o Query, e ele vai retornar linhas da tabela

	linhas, erro := db.Query("select *from usuaios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuarios"))
		defer linhas.Close()

		// vou criar um slice de usuarios
		var usuarios []usuario

		for linhas.Next() { //para cade linha que for retornada, o FOR fara um iteração pra gente (linha 81)
			var usuario usuario //estou colocando dados no meus slice usuarios

			if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
				w.Write([]byte("Erro ao escanear o usuario"))
				return
			}

			usuarios = append(usuarios, usuario)
		}

		w.WriteHeader(http.StatusOK)
		//agora vou transformar esse meu slice de usuario em um json
		if erro := json.NewEncoder(w).Encode(usuarios); erro != nil { //new enconder irá codificar nossos dados para um json, e os dados nós passamos dessa forma
			w.Write([]byte("Erro ao converter os usuarios para JSON"))
			return
		}
	}
}

// BuscarUsuario : traz um usuário específico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r) //irá me retornar todos os parametros

	//especificando o parametro que eu quero retornar

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32) //convertendo meu parametro ID para inteiro. é decimal e depois o tamanho do uint
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	//Vou abrir o banco

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuario!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil { //vou scannear a linha e jogar pra dentro do usuario
			w.Write([]byte("Erro ao escanar o usuario!"))
		}
	}
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil { //new enconder irá codificar nossos dados para um json, e os dados nós passamos dessa forma
		w.Write([]byte("Erro ao conectar o usuario para JSON!"))
		return
	}
}

// AtualizarUsuario altera os dados de um usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r) //irá me retornar todos os parametros
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro ID"))
		return
	}

	//primeiro vou ler o corpo da requisição pra depois abrir requisição com o banco
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body) //ioutil é muito usado pra fazer leitura de entrada e saída de dados. meu readAll vai ler o meu corpo que está dentro do "r" da assinatura
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct!"))
		return
	}

	//agora ja posso abrir a conexao!
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!")) //criamos um comando de inserção que é muito utilizado pra evitar o ataque chamado "sql injection"
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuario!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario deleta um usuario do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r) //irá me retornar todos os parametros
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro so converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?") // NÃO ESQUECE DO WHERE, senão você irá apagar TODOS OS USUARIOS!
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuario!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
