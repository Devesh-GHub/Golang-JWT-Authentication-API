package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTUser struct {
	ID           primitive.ObjectID `bson:"_id"`
	FullName     string             `json:"full_name" validate:"required,min=3,max=100"`
	Password     string             `json:"password" validate:"required,min=6"`
	Email        string             `json:"email"  validate:"email,required"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}
