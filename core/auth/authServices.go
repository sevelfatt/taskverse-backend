package auth

import (
	"context"
	"errors"
	"time"

	"github.com/sevelfatt/taskverse-backend/lib"
	"github.com/sevelfatt/taskverse-backend/models"
	"github.com/sevelfatt/taskverse-backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterService(username string, email string, password string) (string, error) {
	db := lib.MongoClient.Database("taskverse")

	userCollection := db.Collection("users")

	var existingUser bson.M
	err := userCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err == nil {
		return "", errors.New("User already exists")
	}

	currentTime := time.Now()

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := models.User{
		UUID:       utils.GenerateUUID(),
		Username:   username,
		Email:      email,
		Password:   hashedPassword,
		IsVerified: false,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	_, err = userCollection.InsertOne(
		context.TODO(),
		user,
	)

	if err != nil {
		return "", err
	}

	return "User registered successfully", nil

}

func LoginService(email string, password string) (string, error) {
	db := lib.MongoClient.Database("taskverse")

	userCollection := db.Collection("users")

	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err != nil {
		return "", errors.New("User not found")
	}

	if utils.VerifyPassword(password, existingUser.Password) != nil {
		return "", errors.New("Invalid password")
	}

	token, err := utils.CreateJwtToken(existingUser.UUID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func GetUserService(userUUID string) (models.User, error) {
	db := lib.MongoClient.Database("taskverse")

	userCollection := db.Collection("users")


	var existingUser models.User

	err := userCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "uuid", Value: userUUID}}).Decode(&existingUser)
	if err != nil {
		return models.User{}, errors.New("User not found")
	}

	return existingUser, nil
}