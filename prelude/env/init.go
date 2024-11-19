package env

import "github.com/joho/godotenv"

func InitDotenv() {
	if err := godotenv.Load(); err != nil {
		panic("env file is not found")
	}
}
