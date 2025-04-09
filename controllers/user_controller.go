package controllers

import (
	"encoding/json"
	"microservicio/models"
	"microservicio/services"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	result, err := services.CreateUserService(user)
	if err != nil {
		http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := services.GetUserByIDService(id)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsersService()
	if err != nil {
		http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	err := services.UpdateUserService(id, user)
	if err != nil {
		http.Error(w, "Error al actualizar usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := services.DeleteUserService(id)
	if err != nil {
		http.Error(w, "Error al eliminar usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
