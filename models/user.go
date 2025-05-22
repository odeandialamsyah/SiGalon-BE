package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	RoleID   primitive.ObjectID `json:"role_id" bson:"role_id"`
	Role     Role               `json:"role" bson:"role"`
}

type Role struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64           `json:"price" bson:"price"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	ImageURL    string             `json:"image_url" bson:"image_url"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	User        User               `json:"user" bson:"user"`
}

