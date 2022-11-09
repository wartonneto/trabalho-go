package models

import "github.com/wartonneto/trabalho-go/db"

func Update(id int64, livro Livro) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE livros SET title=$2, description=$3, value=$4 WHERE id=$1`, id, livro.Title, livro.Description, livro.Value)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
