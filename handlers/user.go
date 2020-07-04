package handlers

import (
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"context"
	"encoding/json"
	"net/http"
)

func CreateEmployee(svc models.Dao) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err,
				Message: "Error decoding request body",
			}, http.StatusBadRequest)
			return
		}

		err = svc.CreateUser(context.TODO(), user)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err,
				Message: "Error querying mongo",
			}, http.StatusBadRequest)
			return
		}

		utility.NewJSONWriter(w).Write(models.Response{
			Error:   nil,
			Message: "User Created Successsfully",
		}, http.StatusBadRequest)
	}
}
