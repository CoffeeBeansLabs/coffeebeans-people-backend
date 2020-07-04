package handlers

import (
	"coffeebeans-people-backend/auth"
	"coffeebeans-people-backend/models"
	"net/http"
)

func Login(svc models.Dao, authSvc auth.AuthSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var user models.User

	}
}
