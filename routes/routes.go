package routes

import (
	"microservicio/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return router
}
