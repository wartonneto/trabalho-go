package models

import (
	"github.com/wartonneto/trabalho-go/db"
)

func Insert(livro Livro) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO livros (title,description, value) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, livro.Title, livro.Description, livro.Value).Scan(&id)

	return
}
