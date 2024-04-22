package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Abrir uma conexão com o banco de dados MySQL
	db, err := sql.Open("mysql", "meu_app_user:meu_app_password@tcp(meubanco.local:3306)/meu_app_db")
	if err != nil {
		fmt.Println("Erro ao abrir conexão com o banco de dados:", err)
		return
	}
	defer db.Close()

	// Testar a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		fmt.Println("Erro ao testar conexão com o banco de dados:", err)
		return
	}

	fmt.Println("Conexão bem-sucedida ao banco de dados!")
}
