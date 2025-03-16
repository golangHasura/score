package handler

import (
	"github.com/gin-gonic/gin"
	"score/src/handler/routes"
)

func getHttpHandler() *Handler {
	return &Handler{
		FormatRoutes: routes.GetFormatRoutes(),
	}
}

type Handler struct {
	FormatRoutes *routes.FormatRoutes
}

func GetTournamentRoutes(c *gin.Engine) {
	c.GET("/tournaments", routes.GetTournaments)
}

func GetFormatRoutes(c *gin.Engine) {
	handler := getHttpHandler()
	c.GET("/formats", handler.FormatRoutes.GetFormats)
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	GetFormatRoutes(r)
	GetTournamentRoutes(r)
	return r
}
