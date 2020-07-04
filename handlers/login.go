package handlers

import (
	"coffeebeans-people-backend/auth"
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"context"
	"encoding/json"
	"net/http"
)

func Login(svc models.ApiSvc, authSvc auth.AuthSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.User
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err,
				Message: "Error decoding request body",
			}, http.StatusBadRequest)
			return
		}

		user, err := svc.LoginUser(context.TODO(), body.Email, body.Password)
		tokenId, err := authSvc.GenerateToken(&user)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err,
				Message: "Error generating token id",
			}, http.StatusBadRequest)
			return
		}

		loginResponse := models.LoginResponse{
			Email:      user.Email,
			Name:       user.Name,
			EmployeeId: user.EmployeeId,
			Role:       user.Role,
			TokenId:    tokenId,
		}

		utility.NewJSONWriter(w).Write(loginResponse, http.StatusOK)

	}
}
