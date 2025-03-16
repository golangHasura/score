package handler

import (
	"github.com/gin-gonic/gin"
	"score/src/handler/routes"
)

func getHttpHandler() *Handler {
	return &Handler{
		FormatRoutes:     routes.GetFormatRoutes(),
		TournamentRoutes: routes.GetTournamentRoutes(),
	}
}

type Handler struct {
	FormatRoutes     *routes.FormatRoutes
	TournamentRoutes *routes.TournamentRoutes
}

func GetTournamentRoutes(c *gin.Engine) {
	handler := getHttpHandler()
	c.GET("/tournaments", handler.TournamentRoutes.GetTournaments)
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
