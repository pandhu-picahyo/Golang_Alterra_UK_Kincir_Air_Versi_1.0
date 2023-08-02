package main

import (
	"os"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/route"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	loadEnv()
	config.Connect_Database()
	e_echo := echo.New()
	e_echo = route.Inisiasi_Rute(e_echo)
	e_echo.Start(getPort())
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error dalam loading file .env")
	}
}
