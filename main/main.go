package main

import (
	"Blast/config/database"
	"Blast/internal/user/handler"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userHandler := handler.NewUserHandler(db)

	http.HandleFunc("/users", userHandler.HandleUsers)

	log.Println("User service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

	//conn := database.ConnectDB()
	//defer conn.Close()
	//
	//log.Println("Testing database connection...")
}
