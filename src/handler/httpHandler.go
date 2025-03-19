package handler

import (
	"github.com/gin-gonic/gin"
	"score/src/handler/middleware"
	"score/src/handler/routes"
)

func getHttpHandler() *Handler {
	return &Handler{
		FormatRoutes:     routes.GetFormatRoutes(),
		TournamentRoutes: routes.GetTournamentRoutes(),
		TeamRoutes:       routes.GetTeamRoutes(),
	}
}

type Handler struct {
	FormatRoutes     *routes.FormatRoutes
	TournamentRoutes *routes.TournamentRoutes
	TeamRoutes       *routes.TeamRoutes
}

func GetTournamentRoutes(c *gin.Engine) {
	handler := getHttpHandler()
	c.GET("/tournaments", handler.TournamentRoutes.GetTournaments)
}

func GetFormatRoutes(c *gin.Engine) {
	handler := getHttpHandler()
	c.GET("/formats", handler.FormatRoutes.GetFormats)
}

func GetTeamRoutes(c *gin.Engine) {
	handler := getHttpHandler()
	c.GET("/teams", handler.TeamRoutes.GetTeam)
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Use(middleware.InsertTables())
	GetFormatRoutes(r)
	GetTournamentRoutes(r)
	GetTeamRoutes(r)
	return r
}
