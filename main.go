package main

import (
	"os"
	"log"
	"fmt"

	"github.com/gin-gonic/gin"
	"ali_openapi/ali_auth"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	ali_auth.Setup(r.Group("/"))

	port := "9000"
    if envPort := os.Getenv("PORT"); envPort != "" {
        port = envPort
    }
    addr := fmt.Sprintf("127.0.0.1:%s", port)
    r.Run(addr)
}
