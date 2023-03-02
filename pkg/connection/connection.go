package connection

import (
	"fmt"
	"go-bootcamp/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:191491@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
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
	// CreateDatabase()
	Migrate()
	return DB
}

func CreateDatabase() {
	DB.Migrator().DropTable(&models.Book{})
	DB.Migrator().CreateTable(&models.Book{})
}

func Migrate() {
	DB.Migrator().AutoMigrate(&models.Book{})
}
