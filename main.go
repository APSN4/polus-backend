package main

import (
	"github.com/joho/godotenv"
	"os"
	"polus-backend/app/router"
	"polus-backend/config"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
