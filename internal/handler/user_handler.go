package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bkgediya/go_rss_aggregator/internal/db"
	"github.com/Bkgediya/go_rss_aggregator/internal/model"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// decode incoming json to go struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// call the database to create user
	id, err := db.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// return created user as a json object
	user.ID = id
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// call the database to create user
	user, err := db.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// return created user as a json object
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}
