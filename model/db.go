package model

import (
	"database/sql"
)

// Inicializa-se a conexão com o banco de dados

func ConnectDB() (connection *sql.DB) {

	connection, err := sql.Open(Driver, DNS)
	if err != nil {
		panic(err.Error())
	}
	return connection

}
