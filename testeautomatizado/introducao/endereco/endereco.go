package enderecos

//Dentro de testeautomatizados, eu dou o comando "go mod init testeautomatizados-testes" no terminal!

import "strings"

// TipoDeEndereco verifiva se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"avenida", "rua", "estrada", "rodovia", "viela", "beco", "travessa"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco) //passa meu endereço pra letra minuscula

	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	/*Transforma minha string num slice, separando as palavras toda vez que houver um espaço
	e me retorna o primeiro item desse slice*/

	/*Agora vou ver se a minha primeira palavra do endereço, que indica o longradouro
	é válido!*/

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo inválido"
}
