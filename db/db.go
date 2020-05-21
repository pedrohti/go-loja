package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pedrohti/go-loja/errors"
)

const (
	driver   = "postgre"
	port     = 5432
	user     = "postgres"
	dbname   = "alura_loja"
	password = "123456"
	host     = "localhost"
	ssl      = "disable"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, ssl)

	db, err := sql.Open(driver, conexao)
	errors.CheckErrorMsg(err, "Erro ao conectar com o banco")

	return db
}
