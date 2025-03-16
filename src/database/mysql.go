package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	DBOp *gorm.DB
}

// ConnectSql connects to the db and save the value in Pointer struct
func ConnectSql(dbString string) (dbConn *DB, err error) {
	dbConn = &DB{}
	database, err := gorm.Open(mysql.Open(dbString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return dbConn, err
	}
	dbConn.DBOp = database
	return dbConn, nil
}
