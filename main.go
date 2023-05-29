package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	metloGin "github.com/metlo-labs/metlo/ingestors/golang/gin"
	"github.com/metlo-labs/metlo/ingestors/golang/metlo"
)

func main() {
	r := gin.Default()
	metlo := metlo.InitMetlo("https://app.metlo.com", os.Getenv("METLO_API_KEY"))
	instrumentation := metloGin.Init(metlo)
	r.Use(instrumentation.Middleware)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:3031")
}
