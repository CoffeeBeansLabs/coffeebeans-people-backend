package api

import (
	"coffeebeans-people-backend/models"
	"context"
)

type ApiSvc struct {
	DbSvc       models.Dao
}

func (apiSvc *ApiSvc) RegisterUser(ctx context.Context, user models.User) error {

	err := apiSvc.DbSvc.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
