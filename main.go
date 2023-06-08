package main

import (
	"log"

	"github.com/Pastafarian-maiden/pet-API/configs"
	r "github.com/Pastafarian-maiden/pet-API/internal"
)

func main() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r.Run()
}
