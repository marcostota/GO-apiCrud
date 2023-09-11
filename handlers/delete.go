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

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Errror converting id to int %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar registro %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram deletados %v registros", rows)

	}
	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Registro removido com sucesso %d", id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
