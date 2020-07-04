package handlers

import (
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"net/http"
)

func UpdateProfile(apiSvc models.ApiSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(constants.USER_KEY).(models.User)
		if len(user.Email) < 1 {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   "Unauthorized",
				Message: "Invalid token",
			}, http.StatusUnauthorized)
			return
		}


		utility.NewJSONWriter(w).Write(user, http.StatusOK)

	}
}
