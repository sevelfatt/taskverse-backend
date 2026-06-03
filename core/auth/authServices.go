package auth

import (
	"context"
	"errors"
	"time"

	"github.com/sevelfatt/taskverse-backend/lib"
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

	_, err = userCollection.InsertOne(
		context.TODO(),
		map[string]any{
			"uuid"			: 	utils.GenerateUUID(),
			"username"		:   username,
			"email"			:	email,
			"password"		:   password,
			"created_at"	: 	currentTime,
			"updated_at"	: 	currentTime,
			"is_verified"	:	false,
			
		},
	)

	if err != nil {
		return "", err
	}

	return "User registered successfully", nil

}

func LoginService(email string, password string) (bson.M, error) {
	db := lib.MongoClient.Database("taskverse")

	userCollection := db.Collection("users")

	var existingUser bson.M
	err := userCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err != nil {
		return nil, errors.New("User not found")
	}

	if existingUser["password"] != password {
		return nil, errors.New("Invalid password")
	}

	return existingUser, nil

}

func GetUserService(userUUID string) (bson.M, error) {
	db := lib.MongoClient.Database("taskverse")

	userCollection := db.Collection("users")

	var existingUser bson.M
	err := userCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "uuid", Value: userUUID}}).Decode(&existingUser)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return existingUser, nil
}