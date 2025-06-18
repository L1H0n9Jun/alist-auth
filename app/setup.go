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
	aliOpen.Any("/token", aliAccessToken)
	aliOpen.Any("/refresh", aliAccessToken)
	aliOpen.Any("/code", aliAccessToken)
	aliOpen.Any("/qr", aliQrcode)
	aliOpen.Any("/xbcloud", xbAccessToken)
}
