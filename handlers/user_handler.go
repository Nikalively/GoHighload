package handlers

import (
	"encoding/json"
	"gohighload/models"
	"gohighload/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var userService = services.NewUserService()

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := userService.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, exists := userService.GetByID(id)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	savedUser := userService.Create(user)
	go services.LogUserAction("CREATE", savedUser.ID)
	go services.SendNotification(savedUser, "created")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(savedUser)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedUser, exists := userService.Update(id, user)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	go services.LogUserAction("UPDATE", id)
	go services.SendNotification(updatedUser, "updated")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, exists := userService.GetByID(id)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if !userService.Delete(id) {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	go services.LogUserAction("DELETE", id)
	go services.SendNotification(user, "deleted")
	w.WriteHeader(http.StatusNoContent)
}
