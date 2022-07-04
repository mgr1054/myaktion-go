package db

import (
	"fmt"
	"os"

	"github.com/mgr1054/myaktion-go/src/myaktion/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// init DB
	var err error
	dsn := fmt.Sprintf("root:root@tcp(%s)/myaktion?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_CONNECT"))
	log.Info("Using DSN for DB:", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	log.Info("Starting automatic migrations")
	if err := DB.Debug().AutoMigrate(&model.Campaign{}); err != nil {
		panic(err)
	}
	if err := DB.Debug().AutoMigrate(&model.Donation{}); err != nil {
		panic(err)
	}
	log.Info("Automatic migrations finished")
}
