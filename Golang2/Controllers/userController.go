package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	database "github.com/devesh/mongoapi/Database"
	helper "github.com/devesh/mongoapi/Helpers"
	model "github.com/devesh/mongoapi/Models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection = database.OpenCollection(database.Client, "user")

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	count, err := userCollection.CountDocuments(context.Background(), bson.M{"email": user.Email})
	if err != nil {
		http.Error(w, "Error checking email existence", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	user.Password, err = helper.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.FullName, "USER", user.User_id)
	user.Token = &token
	user.RefreshToken = &refreshToken

	_, insertErr := userCollection.InsertOne(context.Background(), user)
	if insertErr != nil {
		http.Error(w, "User not created", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var foundUser model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	passwordValid := helper.CheckPasswordHash(user.Password, foundUser.Password)
	if !passwordValid {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.FullName, "USER", foundUser.User_id)
	helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	json.NewEncoder(w).Encode(map[string]string{"token": token, "refresh_token": refreshToken})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	var user model.User
	err := userCollection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
