package routes

import(
	"github.com/gorilla/mux"
	"BetterGOChat/controllers"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/api/messages", controllers.GetAllMessages).Methods("GET")
	router.HandleFunc("/api/messages", controllers.PostMessage).Methods("POST")
	router.HandleFunc("/api/messages/{id}", controllers.DeleteMessage).Methods("DELETE")
}