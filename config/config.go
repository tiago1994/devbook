package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnection = ""
	Port             = 0
)

func LoadConfig() {
	var loadError error
	if loadError = godotenv.Load(); loadError != nil {
		log.Fatal(loadError)
	}

	Port, loadError = strconv.Atoi(os.Getenv("API_PORT"))
	if loadError != nil {
		Port = 9000
	}

	StringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
