package repositories

import ()

var (
	userCollection *mongo.Collection
	roleCollection *mongo.Collection
)

func InitCollections() {
	userCollection = config.DB.Collection("users")
	roleCollection = config.DB.Collection("roles")
}