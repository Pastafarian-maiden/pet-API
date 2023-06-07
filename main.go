package main

import (
	"log"

	r "github.com/Pastafarian-maiden/pet-API/internal"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	r.Run()

}

// // Simple helper function to read an environment or return a default value
// func getEnv(key string, defaultVal string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return defaultVal
// }
