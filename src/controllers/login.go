package controllers

import (
	"devbook/src/authentication"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repository"
	"devbook/src/responses"
	"devbook/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request, requestError := ioutil.ReadAll(r.Body)
	if requestError != nil {
		responses.Error(w, http.StatusUnprocessableEntity, requestError)
		return
	}

	var user models.User
	if requestError = json.Unmarshal(request, &user); requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	db, requestError := database.Connect()
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	userDatabase, requestError := repository.GetByEmail(user.Email)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	if requestError = security.CheckPassword(userDatabase.Password, user.Password); requestError != nil {
		responses.Error(w, http.StatusUnauthorized, requestError)
		return
	}

	token, requestError := authentication.CreateToken(userDatabase.ID)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}
	w.Write([]byte(token))
}
