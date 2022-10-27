package models

import "github.com/wartonneto/trabalho-go/db"

func GetAll() (livros []Livro, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM livros`)

	if err != nil {
		return
	}

	for rows.Next() {
		var livro Livro

		err = rows.Scan(&livro.ID, &livro.Title, &livro.Description, &livro.Value)
		if err != nil {
			continue
		}

		livros = append(livros, livro)

	}
	return
}
