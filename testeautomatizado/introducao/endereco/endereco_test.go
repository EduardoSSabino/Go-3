package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	//Essa é a assinatura de uma função de teste
	//começa obrigatóriamente com a palavra Test + nome da função que irá ser testada
	//e depois recebe como parametro "t *testing.T" por PADRÃO!

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Imigrantes", "Rodovia"},
		{"Praças das rosas", "Tipo inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado")
		}
	}

}
