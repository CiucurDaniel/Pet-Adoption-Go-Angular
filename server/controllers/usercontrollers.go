package controllers

import (
	"ciucurdaniel.com/pet-adoption-server/database"
	"ciucurdaniel.com/pet-adoption-server/entities"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// CreateUser POST - user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func checkIsUserIdExists(userId string) bool {
	var user entities.User
	database.Instance.First(&user, userId)
	if user.ID == 0 {
		return false
	}
	return true
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	if checkIsUserIdExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found")
		return
	}

	var user entities.User
	database.Instance.First(&user, userId)
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users[] entities.User
	database.Instance.Find(&users)
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	if checkIsUserIdExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	}

	var user entities.User
	database.Instance.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Save(&user)
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	if checkIsUserIdExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	}

	var user entities.User
	database.Instance.Delete(&user, userId)
}