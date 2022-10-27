package models

import "github.com/wartonneto/trabalho-go/db"

func Get(id int64) (livro Livro, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM livros WHERE id=$1`, id)

	err = row.Scan(&livro.ID, &livro.Title, &livro.Description, &livro.Value)

	return
}
