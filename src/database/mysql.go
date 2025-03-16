package database

import (
	"fmt"
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
	fmt.Println("db string is:", dbString)
	database, err := gorm.Open(mysql.Open(dbString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return dbConn, err
	}
	fmt.Println("databse", database)
	dbConn.DBOp = database
	return dbConn, nil
}
