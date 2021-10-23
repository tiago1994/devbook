package main

import (
	"devbook/config"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	fmt.Printf("Running on port %d", config.Port)
	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
