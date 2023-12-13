package main

import (
	"github.com/joho/godotenv"
	"github.com/sigchat/sc-users/pkg"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env files")
	}
	pkg.InitAPI()
}
