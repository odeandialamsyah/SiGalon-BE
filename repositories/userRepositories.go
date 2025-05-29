package repositories

import (
	"context"
	"sigalon-be/config"
	"sigalon-be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection *mongo.Collection
	roleCollection *mongo.Collection
)

func InitCollections() {
	userCollection = config.DB.Collection("users")
	roleCollection = config.DB.Collection("roles")
}

// Fungsi untuk mendapatkan user berdasarkan username
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	// Get role information
	var role models.Role
	err = roleCollection.FindOne(context.Background(), bson.M{"_id": user.RoleID}).Decode(&role)
	if err != nil {
		return models.User{}, err
	}

	user.Role = role
	return user, nil
}

// Fungsi untuk mendapatkan user berdasarkan email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	// Get role information
	var role models.Role
	err = roleCollection.FindOne(context.Background(), bson.M{"_id": user.RoleID}).Decode(&role)
	if err != nil {
		return models.User{}, err
	}

	user.Role = role
	return user, nil
}

// Fungsi untuk menambahkan user baru
func CreateUser(user models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
