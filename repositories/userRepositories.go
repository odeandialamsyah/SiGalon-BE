package repositories

import (
	"context"
	"sigalon-be/config"
	"sigalon-be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetAllUsers returns all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	// Get role information for each user
	for i := range users {
		role, err := GetRoleByID(users[i].RoleID.Hex())
		if err != nil {
			continue
		}
		users[i].Role = role
	}

	return users, nil
}

// GetUserByID returns a user by ID
func GetUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	// Get role information
	role, err := GetRoleByID(user.RoleID.Hex())
	if err != nil {
		return models.User{}, err
	}
	user.Role = role

	return user, nil
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

// Fungsi untuk mendapatkan role berdasarkan ID
func GetRoleByID(roleID string) (models.Role, error) {
	var role models.Role
	err := roleCollection.FindOne(context.Background(), bson.M{"_id": roleID}).Decode(&role)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

// Fungsi untuk menambahkan user baru
func CreateUser(user models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser updates a user's information
func UpdateUser(user models.User) error {
	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"email":    user.Email,
		},
	}

	if !user.RoleID.IsZero() {
		update["$set"].(bson.M)["role_id"] = user.RoleID
	}

	_, err := userCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		update,
	)
	return err
}

// DeleteUser deletes a user
func DeleteUser(id primitive.ObjectID) error {
	_, err := userCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
