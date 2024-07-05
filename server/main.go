package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var tag = os.Getenv("TAG")

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		slog.Info("server running", slog.String("tag", tag))
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"tag":     tag,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
