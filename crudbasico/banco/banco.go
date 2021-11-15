package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conxão com o MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) { //me retorn aum sql.DB que será meu banco e me retorna um erro também
	stringConexao := "eduardo:@Sabiedusilvawheel13@/bancodedaos?charset=utf8&parseTime=True&loc=local"

	db, erro := sql.Open("mysql", stringConexao) //abre meu banco de dados
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
