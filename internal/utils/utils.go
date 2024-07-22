package utils

import "github.com/joho/godotenv"

func CheckEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Error while loading .env!")
	}
}
