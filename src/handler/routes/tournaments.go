package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tournament struct {
	Name       string `json:"name"`
	Format     string `json:"format"`
	LeagueName string `json:"league"`
}

func GetTournaments(c *gin.Context) {

	tournaments := []Tournament{
		{
			Name:       "IPL",
			Format:     "T20",
			LeagueName: "Domestic",
		},
		{
			Name:       "World Cup",
			Format:     "ODI",
			LeagueName: "International",
		},
	}
	c.IndentedJSON(http.StatusOK, tournaments)
}
