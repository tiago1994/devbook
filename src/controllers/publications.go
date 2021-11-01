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

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, requestError := authentication.GetUserId(r)
	if requestError != nil {
		responses.Error(w, http.StatusUnauthorized, requestError)
		return
	}

	request, requestError := ioutil.ReadAll(r.Body)
	if requestError != nil {
		responses.Error(w, http.StatusUnprocessableEntity, requestError)
		return
	}

	var publication models.Publication
	if requestError = json.Unmarshal(request, &publication); requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	publication.CreatedID = userIDToken

	if requestError = publication.Preparate(); requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	db, requestError := database.Connect()
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}
	defer db.Close()

	repository := repository.NewPublicationRepository(db)
	publication.ID, requestError = repository.Create(publication)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

func GetAllPublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, requestError := authentication.GetUserId(r)
	if requestError != nil {
		responses.Error(w, http.StatusUnauthorized, requestError)
		return
	}

	db, requestError := database.Connect()
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}
	defer db.Close()

	repository := repository.NewPublicationRepository(db)
	publications, requestError := repository.GetAll(userIDToken)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

func GetPublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, requestError := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewPublicationRepository(db)
	user, requestError := repository.GetBy(publicationID)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, requestError := authentication.GetUserId(r)
	if requestError != nil {
		responses.Error(w, http.StatusUnauthorized, requestError)
		return
	}

	parameters := mux.Vars(r)
	publicationID, requestError := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewPublicationRepository(db)
	publicacationDB, requestError := repository.GetBy(userIDToken)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	if publicacationDB.CreatedID != userIDToken {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden operation"))
		return
	}

	request, requestError := ioutil.ReadAll(r.Body)
	if requestError != nil {
		responses.Error(w, http.StatusUnprocessableEntity, requestError)
		return
	}

	var publication models.Publication
	if requestError = json.Unmarshal(request, &publication); requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	if requestError = publication.Preparate(); requestError != nil {
		responses.Error(w, http.StatusBadRequest, requestError)
		return
	}

	if requestError = repository.UpdatePublication(publicationID, publication); requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, requestError := authentication.GetUserId(r)
	if requestError != nil {
		responses.Error(w, http.StatusUnauthorized, requestError)
		return
	}

	parameters := mux.Vars(r)
	publicationID, requestError := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewPublicationRepository(db)
	publicacationDB, requestError := repository.GetBy(userIDToken)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	if publicacationDB.CreatedID != userIDToken {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden operation"))
		return
	}

	if requestError = repository.DeletePublication(publicationID); requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetPublicationsByUser(w http.ResponseWriter, r *http.Request) {
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

	repository := repository.NewPublicationRepository(db)
	publications, requestError := repository.GetByUser(userID)
	if requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

func Like(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, requestError := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewPublicationRepository(db)
	if requestError = repository.Like(publicationID); requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func Dislike(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, requestError := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repository := repository.NewPublicationRepository(db)
	if requestError = repository.Dislike(publicationID); requestError != nil {
		responses.Error(w, http.StatusInternalServerError, requestError)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
