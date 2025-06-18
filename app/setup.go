package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

func initVar() {
	// client
	aliOpenClientID = os.Getenv("ALI_OPEN_CLIENT_ID")
	aliOpenClientSecret = os.Getenv("ALI_OPEN_CLIENT_SECRET")
}

func Setup(g *gin.RouterGroup) {
	initVar()
	aliOpen := g.Group("/ali_open")
	aliOpen.POST("/token", aliAccessToken)
	aliOpen.POST("/refresh", aliAccessToken)
	aliOpen.POST("/code", aliAccessToken)
	aliOpen.POST("/qr", aliQrcode)
	aliOpen.POST("/xbcloud", xbAccessToken)
}
