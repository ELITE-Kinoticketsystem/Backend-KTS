package main

import (
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/joho/godotenv"
)

func main() {
	// Load Environment variables
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file:: %v", err)
	}
	log.Println("Environment variables loaded successfully")

	// Initialize Database
	dbConnection, _ := managers.InitializeDB()
	//log.Print(dbConnection)

	errPing := dbConnection.Ping()
	if errPing != nil {
		log.Printf("Error while pinging database: %s", errPing.Error())
	} else {
		log.Printf("Successfully connected to database")
	}
}
