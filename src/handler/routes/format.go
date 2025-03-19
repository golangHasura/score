package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score/src/internal/dto"
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

	fetchedFormats, err := f.FormatRepo.FetchFormats(db)
	if err != nil {
		errRes := map[string]string{
			"code":    h.CommonErrorCode,
			"message": err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, errRes)
		return
	}
	var formats []dto.Format
	for _, fetchedFormat := range fetchedFormats {
		format := dto.Format{
			EntityId: fetchedFormat.EntityId,
			Name:     fetchedFormat.Name,
		}
		formats = append(formats, format)
	}
	res := dto.RetrieveAllResponse{Total: len(formats), Data: formats}
	c.IndentedJSON(http.StatusOK, res)
}
