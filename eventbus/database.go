package eventbus

import (
	"os"

	"github.com/jinzhu/gorm"

	// the Poestgres dialects is required for connecting to PG databases
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("DATABASE_HOST")+
			" user="+os.Getenv("DATABASE_USER")+
			" dbname="+os.Getenv("DATABASE_NAME")+
			" password="+os.Getenv("DATABASE_PASSWORD")+
			" sslmode=disable"+
			" connect_timeout=5")

	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Subscription{})
}
