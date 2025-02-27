package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	helper "github.com/devesh/mongoapi/Helpers"
	model "github.com/devesh/mongoapi/Models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Devesh:Devesh%402004@cluster0.kjzfj.mongodb.net/?retryWrites=true&w=majority"
const dbName = "hapeemail"
const colName = "UserData"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Successful")

	collection = client.Database(dbName).Collection(colName)

	indexModels := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.Background(), indexModels)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected and unique index created")
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = InsertOneUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var foundUser model.User
	err = collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !helper.CheckPasswordHash(user.Password, foundUser.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := helper.GenerateJWT(foundUser.Email, foundUser.User_id)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func InsertOneUser(user model.User) error {
	if user.Email == "" || user.Password == "" || user.FullName == "" {
		return fmt.Errorf("all fields are required")
	}

	var existingUser model.User
	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return fmt.Errorf("email already exists")
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	user.Password = hashPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	fmt.Println("Inserted one user details with id:", inserted.InsertedID)
	return nil
}

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/api/register", RegisterUserHandler).Methods("POST")
	router.HandleFunc("/api/login", LoginUserHandler).Methods("POST")
}
