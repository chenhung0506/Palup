package handlers

import (
	"encoding/json"
	"net/http"
	"playsee/models"
)

func Test1(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Array []interface{} `json:"Array"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request: Unable to parse JSON", http.StatusBadRequest)
		return
	}

	head := models.CreateLinkedList(requestBody.Array)

	models.PrintLinkedList(head)

	response, err := json.Marshal(head)
	if err != nil {
		http.Error(w, "Internal Server Error: Unable to serialize linked list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
