package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marcostota/apicrud/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting id to int %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Error ao buscar o todo %v ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}
