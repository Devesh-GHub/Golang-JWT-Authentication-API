package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FullName    string             `bson:"full_name" json:"full_name" validate:"required,min=3,max=100"`
	Email       string             `bson:"email" json:"email" validate:"email,required"`
	Password    string             `bson:"password" json:"password" validate:"required,min=6"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number" validate:"required"`
	UserType    string             `bson:"user_type" json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	UserID      string             `bson:"user_id" json:"user_id"`
}
