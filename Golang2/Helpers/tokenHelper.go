package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

var userCollection *mongo.Collection
var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, fullName string, userType string, userID string) (string, string, error) {
	claims := &SignedDetails{
		Email:    email,
		FullName: fullName,
		UserID:   userID,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(168 * time.Hour).Unix(),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateObj := bson.D{
		{"token", signedToken},
		{"refresh_token", signedRefreshToken},
		{"updated_at", time.Now()},
	}

	filter := bson.M{"user_id": userID}
	opt := options.Update().SetUpsert(true)

	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, opt)
	if err != nil {
		log.Fatal(err)
	}
}
