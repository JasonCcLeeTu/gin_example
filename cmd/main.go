package main

import (
	handle "gin/internal/delivery"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	engine := gin.Default()

	handle.ApiRouter(engine)

	engine.Run(":8070")
}
