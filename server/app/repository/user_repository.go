package repository

import (
	"billbuddies/server/app/helper"
	"billbuddies/server/app/model"
	"billbuddies/server/app/provider"
	"billbuddies/server/core/logger"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongoDriver.Collection
}

// NewUserRepository initializes a new UserRepository
func NewUserRepository(mongoProvider *provider.MongoProvider) *UserRepository {
	db := mongoProvider.GetDB()
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	// Hash the password before storing it
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		logger.Log.Error("error hashing password", logger.Error(err))
		return err
	}
	user.Password = hashedPassword

	_, err = r.collection.InsertOne(ctx, user)
	if err != nil {
		logger.Log.Error("error creating user", logger.Error(err))
		return err
	}
	return nil
}

func (r *UserRepository) ReadUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongoDriver.ErrNoDocuments {
			return nil, nil
		}
		logger.Log.Error("error reading user by username", logger.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, username string, update bson.M) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": update})
	if err != nil {
		logger.Log.Error("error updating user", logger.Error(err))
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, username string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"username": username})
	if err != nil {
		logger.Log.Error("error deleting user", logger.Error(err))
		return err
	}
	return nil
}