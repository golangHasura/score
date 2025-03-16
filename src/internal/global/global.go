package global

import (
	"os"
	"score/src/database"
)

var db *database.DB

func SetDBToGlobal() (err error) {
	if db, err = database.ConnectSql(os.Getenv("DB_CONNECTION")); err != nil {
		return err
	}
	return nil
}

func GetDbFromGlobal() *database.DB {
	return db
}
