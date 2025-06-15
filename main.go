package main

import (
	"alist-auth/app"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	app.Setup(r.Group("/api"))

	port := "9000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	addr := fmt.Sprintf("127.0.0.1:%s", port)
	r.Run(addr)
}
