package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marcostota/apicrud/models"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting id to int %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding JSON %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro %v ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Messgae": fmt.Sprintf("Todo atualizado com successo : %d", id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
