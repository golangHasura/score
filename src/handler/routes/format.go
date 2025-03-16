package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score/src/internal/global"
	h "score/src/internal/helper"
	"score/src/repo"
	v1Repo "score/src/repo/v1"
)

func GetFormatRoutes() *FormatRoutes {
	return &FormatRoutes{
		FormatRepo: v1Repo.NewFormatRepo(),
	}
}

type FormatRoutes struct {
	FormatRepo repo.FormatRepo
}

func (f *FormatRoutes) GetFormats(c *gin.Context) {
	// Getting DB from global
	db := global.GetDbFromGlobal()

	formats, err := f.FormatRepo.FetchFormats(db)
	if err != nil {
		errRes := map[string]string{
			"code":    h.CommenErrorCode,
			"message": err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, errRes)
	}
	c.IndentedJSON(http.StatusOK, formats)
}
