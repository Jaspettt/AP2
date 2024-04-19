package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBUSER = "jaspet"
	DBPASS = "1337"
	DBNAME = "VinylsGolang"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true",
		DBUSER, DBPASS, DBNAME)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
