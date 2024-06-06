package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func RunServer() {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders: "Origin, Authorization, Content-Type, Options",
		MaxAge:         50 * time.Second,
	}))

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Bienvenido a la api")
	})

	SetupRoutes(server)

	if err := server.Run("localhost:3000"); err != nil {
		slog.Error("Run Server: ", err)
	}
}
