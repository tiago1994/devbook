package controllers

import (
	"devbook/src/authentication"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repository"
	"devbook/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, requestError := ioutil.ReadAll(r.Body)
	if requestError != nil {
		responses.Error(w, http.StatusUnprocessableEntity, requestError)
		return
	}

	var newUser models.User
	if requestError = json.Unmarshal(request, &newUser); requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	if requestError = newUser.Preparate("create"); requestError != nil {
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
	newUser.ID, requestError = repository.Create(newUser)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusCreated, newUser)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, requestError := database.Connect()
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	users, requestError := repository.GetAll(nameOrNick)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, requestError := strconv.ParseUint(parameters["userId"], 10, 64)
	if requestError != nil {
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
	user, requestError := repository.GetBy(userID)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, requestError := strconv.ParseUint(parameters["userId"], 10, 64)
	if requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	userIDToken, requestError := authentication.GetUserId(r)
	if requestError != nil {
		responses.Error(w, http.StatusUnauthorized, requestError)
		return
	}

	if userIDToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden operation"))
		return
	}

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

	if requestError = user.Preparate("edit"); requestError != nil {
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
	if requestError = repository.Update(userID, user); requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, requestError := strconv.ParseUint(parameters["userId"], 10, 64)
	if requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	userIDToken, requestError := authentication.GetUserId(r)
	if requestError != nil {
		responses.Error(w, http.StatusUnauthorized, requestError)
		return
	}

	if userIDToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden operation"))
		return
	}

	db, requestError := database.Connect()
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	if requestError := repository.Delete(userID); requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
