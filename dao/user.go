package dao

import (
	"coffeebeans-people-backend/models"
	"context"
)

func (service *Service) CreateUser(ctx context.Context, user models.User) error {

	c := service.MongoConn.Collection("users")

	_, err := c.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
