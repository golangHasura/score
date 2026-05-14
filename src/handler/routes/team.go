package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score/src/internal/dto"
	g "score/src/internal/global"
	h "score/src/internal/helper"
	"score/src/repo"
	v1 "score/src/repo/v1"
)

func GetTeamRoutes() *TeamRoutes {
	return &TeamRoutes{
		repo: v1.NewTeamRepo(),
	}
}

type TeamRoutes struct {
	repo repo.TeamRepo
}

func (t *TeamRoutes) GetTeam(c *gin.Context) {
	db := g.GetDbFromGlobal()
	if db == nil || db.DBOp == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"code":    h.CommonErrorCode,
			"message": "database connection is not initialized",
		})
		return
	}

	fetchedTeams, err := t.repo.FetchTeams(db)
	if err != nil {
		errRes := map[string]string{
			"code":    h.CommonErrorCode,
			"message": err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, errRes)
		return
	}
	var teams []dto.Team
	for _, fetchedTeam := range fetchedTeams {
		team := dto.Team{
			EntityId: fetchedTeam.EntityId,
			Name:     fetchedTeam.Name,
			Logo:     fetchedTeam.Logo,
			Players:  []dto.Player{},
			Format: dto.Format{
				Name: fetchedTeam.Format.Name,
			},
		}
		teams = append(teams, team)
	}
	res := dto.RetrieveAllResponse{Total: len(teams), Data: teams}
	c.IndentedJSON(http.StatusOK, res)
}
