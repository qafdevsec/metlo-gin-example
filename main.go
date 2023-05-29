package main

import (
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	metloGin "github.com/metlo-labs/metlo/ingestors/golang/gin"
	"github.com/metlo-labs/metlo/ingestors/golang/metlo"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

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
	r.GET("/large-response", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": RandStringBytesMaskImprSrcSB(9000),
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:3031")
}
