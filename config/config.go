package config

import (
	"altastore/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDb() {
	err1 := godotenv.Load("./.env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err2 error
	Db, err2 = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err2 != nil {
		panic(err2)
	}

	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Cart{})
	Db.AutoMigrate(&models.Product{})
	Db.AutoMigrate(&models.Category{})
	Db.AutoMigrate(&models.Transaction{})
	Db.AutoMigrate(&models.TransactionDetail{})
	Db.AutoMigrate(&models.PaymentMethod{})
	Db.AutoMigrate(&models.Admin{})
}

// this config for API testing purpose
func InitDBTest() {
	// Since we invoke this from inside of altastore/controller/controllerxxx,
	// we need to cd to parent directory twice
	if err := godotenv.Load("./../../.env"); err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file. Got this error: %v", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME_TEST"))

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	Db.Migrator().DropTable(&models.User{},&models.Category{}, &models.Product{})
	Db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})
}