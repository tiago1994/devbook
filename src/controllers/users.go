package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, requestError := ioutil.ReadAll(r.Body)
	if requestError != nil {
		log.Fatal(requestError)
	}

	var newUser models.User
	if requestError = json.Unmarshal(request, &newUser); requestError != nil {
		log.Fatal(requestError)
	}

	db, requestError := database.Connect()
	if requestError != nil {
		log.Fatal(requestError)
	}

	repository := repository.NewUserRepository(db)
	userId, requestError := repository.Create(newUser)
	if requestError != nil {
		log.Fatal(requestError)
	}

	w.Write([]byte(fmt.Sprintf("User created, id: %d", userId)))
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users!"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
