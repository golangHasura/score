package global

import (
	"fmt"
	"os"
	"score/src/database"
)

var db *database.DB

func SetDBToGlobal() (err error) {
	fmt.Println(os.Getenv("DB_CONNECTION"))
	if db, err = database.ConnectSql(os.Getenv("DB_CONNECTION")); err != nil {
		return err
	}
	return nil
}

func GetDbFromGlobal() *database.DB {
	return db
}
