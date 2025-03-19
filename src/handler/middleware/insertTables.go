package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		for _, tableName := range mapper {
			if !db.DBOp.Migrator().HasTable(tableName) {
				err := db.DBOp.Migrator().AutoMigrate(tableName)
				if err != nil {
					fmt.Println("Auto migrate error:", err)
					return
				}
			}
		}
	}
}
