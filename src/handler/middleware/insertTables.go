package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"score/src/internal/global"
	"score/src/models"
)

var mapper = map[string]interface{}{
	"formats": &models.Format{},
	"teams":   &models.Team{},
}

// InsertTables Inserts tables if it is not present
func InsertTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := global.GetDbFromGlobal()
		if db == nil || db.DBOp == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    "SC01",
				"message": "database connection is not initialized",
			})
			return
		}

		for _, tableModel := range mapper {
			if !db.DBOp.Migrator().HasTable(tableModel) {
				err := db.DBOp.Migrator().AutoMigrate(tableModel)
				if err != nil {
					fmt.Println("Auto migrate error:", err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"code":    "SC01",
						"message": err.Error(),
					})
					return
				}
			}
		}
		c.Next()
	}
}
