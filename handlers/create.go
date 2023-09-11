package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/marcostota/apicrud/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding jsong %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := models.Insert(todo)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Occorreu um erro ao tentar inserir %v", err),
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo inserido com sucesso : %d", id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
