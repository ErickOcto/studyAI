package main

import (
	"log"
	"net/http"
	"studyAI/config"
	"studyAI/controllers/DashboardController"
)

func main() {
	config.InitDatabase()

	//dashboard
	http.HandleFunc("/", DashboardController.Welcome)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
