package repositories

import (
	"context"
	"errors"
	"microservicio/config"
	"microservicio/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return collection.InsertOne(ctx, user)
}

func GetUserByID(id string) (models.User, error) {
	var user models.User
	collection := config.GetCollection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("usuario no encontrado")
		}
		return user, err
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	collection := config.GetCollection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(id string, updated models.User) error {
	collection := config.GetCollection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":     updated.Name,
			"email":    updated.Email,
			"password": updated.Password,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func DeleteUser(id string) error {
	collection := config.GetCollection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
