package controllers

import (
	"go-rest-api-jwt/helpers"
	"go-rest-api-jwt/models"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	userResponse := &models.MyProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}
