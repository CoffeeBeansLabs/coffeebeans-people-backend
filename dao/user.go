package dao

import (
	"coffeebeans-people-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (service *Service) CreateUser(ctx context.Context, user models.User) error {

	c := service.MongoConn.Collection("users")
	model1 := mongo.IndexModel{
		Keys: bson.M{
			"employee_id": user.EmployeeId,
		},
		Options: options.Index().SetUnique(true),
	}

	model2 := mongo.IndexModel{
		Keys: bson.M{
			"email": user.Email,
		},
		Options: options.Index().SetUnique(true),
	}

	c.Indexes().CreateOne(ctx, model1)
	c.Indexes().CreateOne(ctx, model2)

	_, err := c.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) GetUserByEmployeeId(ctx context.Context, employeeId int64) (models.User, error) {
	user := models.User{}

	collection := service.MongoConn.Collection("users")

	doc := collection.FindOne(ctx, bson.M{"employee_id": employeeId})

	err := doc.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *Service) GetUserByCredentials(ctx context.Context, email string, password string) (models.User, error) {
	user := models.User{}

	collection := service.MongoConn.Collection("users")

	doc := collection.FindOne(ctx, bson.M{"email": email, "password": password})

	err := doc.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
