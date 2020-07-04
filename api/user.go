package api

import (
	"coffeebeans-people-backend/models"
	"context"
	"encoding/json"
	"reflect"
)

type ApiSvc struct {
	DbSvc models.Dao
}

func (apiSvc *ApiSvc) RegisterUser(ctx context.Context, user models.User) error {

	err := apiSvc.DbSvc.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (apiSvc *ApiSvc) LoginUser(ctx context.Context, email string, password string) (models.User, bool, error) {
	var isProfileComplete bool
	user, err := apiSvc.DbSvc.GetUserByCredentials(ctx, email, password)
	if err != nil {
		return user, isProfileComplete, err
	}

	isProfileComplete = isProfileCompleted(user)

	return user, isProfileComplete, nil
}

func isProfileCompleted(user models.User) bool {
	var userMandatoryFields models.UserMandatoryFields

	marshalledData, _ := json.Marshal(&user)

	json.Unmarshal(marshalledData, &userMandatoryFields)

	return !reflect.DeepEqual(userMandatoryFields, models.UserMandatoryFields{})

}
