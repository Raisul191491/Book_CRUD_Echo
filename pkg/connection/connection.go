package connection

import (
	"fmt"
	"go-bootcamp/pkg/config"
	"go-bootcamp/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	db := config.LocalConfig

	dsn := fmt.
		Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.DBUser, db.DBPass, db.DBIP, db.DbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("error connecting to DB")
		panic(err)
	}

	fmt.Println("Database Connected")
	DB = d
}

func Migrate() {
	DB.Migrator().AutoMigrate(&models.Book{})
}

func GetDB() *gorm.DB {
	if DB == nil {
		Connect()
	}
	Migrate()
	return DB
}
