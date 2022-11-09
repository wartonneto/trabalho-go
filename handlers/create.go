package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wartonneto/trabalho-go/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var livro models.Livro

	err := json.NewDecoder(r.Body).Decode(&livro)
	if err != nil {
		log.Printf("erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(livro)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir aqui: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Livro inserido com sucesso! id referente: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
