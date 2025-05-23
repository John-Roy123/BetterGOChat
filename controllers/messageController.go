package controllers

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"BetterGOChat/models"
	"BetterGOChat/database"
)

var users []models.Message

func GetAllMessages(w http.ResponseWriter, r *http.Request){
	var message []models.Message
	database.DB.Find(&message)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func PostMessage(w http.ResponseWriter, r *http.Request){
	var message models.Message
	json.NewDecoder(r.Body).Decode(&message)
	database.DB.Create(&message)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var message models.Message
	database.DB.Delete(&message, params["id"])
	w.WriteHeader(http.StatusNoContent)
}