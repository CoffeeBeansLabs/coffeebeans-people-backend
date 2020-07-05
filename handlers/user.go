package handlers

import (
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"context"
	"encoding/json"
	"net/http"
)

func CreateEmployee(apiSvc models.ApiSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err.Error(),
				Message: "Error decoding request body",
			}, http.StatusBadRequest)
			return
		}

		user.Role = "ADMIN" //TODO ACCEPT FROM UI
		err = apiSvc.RegisterUser(context.TODO(), user)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   "Mongo error",
				Message: "Employee Id already exists. Check you employee id",
			}, http.StatusBadRequest)
			return
		}

		utility.NewJSONWriter(w).Write(models.Response{
			Error:   "",
			Message: "User Created Successsfully",
		}, http.StatusOK)
	}
}
