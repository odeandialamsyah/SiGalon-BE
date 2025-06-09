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
	categoryCollection *mongo.Collection
)