package api

import (
	"coffeebeans-people-backend/models"
	"context"
	"encoding/json"
	"fmt"
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

func (apiSvc *ApiSvc) LoginUser(ctx context.Context, email string, password string) (models.User, error) {

	user, err := apiSvc.DbSvc.GetUserByCredentials(ctx, email, password)
	if err != nil {
		return user, err
	}

	isProfileComplete := isProfileComplete(user)
	fmt.Println("COmppllle", isProfileComplete)
	return user, nil
}

func isProfileComplete(user models.User) bool {
	var userMandatoryFields models.UserMandatoryFields

	marshalledData, _ := json.Marshal(&user)

	json.Unmarshal(marshalledData, &userMandatoryFields)

	return reflect.DeepEqual(userMandatoryFields, models.UserMandatoryFields{})

}
