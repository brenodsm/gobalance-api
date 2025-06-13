package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSON envia uma resposta HTTP com status e dados serializados em JSON.
func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	if data == nil {
		w.WriteHeader(statusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Erro ao serializar em JSON: %v\n", err)
	}
}

// WriteErrorJSON envia uma resposta de erro HTTP no formato JSON.
func WriteErrorJSON(w http.ResponseWriter, statusCode int, err error) {
	WriteJSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
