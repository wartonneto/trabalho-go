package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wartonneto/trabalho-go/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	livros, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao obter registros: %v ", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livros)
}
