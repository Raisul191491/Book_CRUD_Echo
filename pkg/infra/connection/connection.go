package connection

import (
	"fmt"
	"go-bootcamp/pkg/infra/config"
	"go-bootcamp/pkg/infra/connection/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config.SetConfig()
	db := config.LocalConfig

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", db.DBUser, db.DBPass, db.DBIP, db.DbName)
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("error connecting to DB")
		panic(err)
	}

	fmt.Println("Database Connected")
	DB = d
}

func GetDB() *gorm.DB {
	if DB == nil {
		Connect()
	}
	CreateDatabase()
	Migrate()
	return DB
}

func CreateDatabase() {
	DB.Migrator().DropTable(&model.Book{})
	DB.Migrator().CreateTable(&model.Book{})
}

func Migrate() {
	DB.Migrator().AutoMigrate(&model.Book{})
}
