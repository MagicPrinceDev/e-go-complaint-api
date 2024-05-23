package mysql

import (
	"e-complaint-api/drivers/mysql/seeder"
	"e-complaint-api/entities"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func ConnectDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Migration(db)
	Seeder(db)

	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(entities.Admin{})
	db.AutoMigrate(entities.User{})
	db.AutoMigrate(entities.Category{})
}

func Seeder(db *gorm.DB) {
	seeder.SeedAdmin(db)
	seeder.SeedCategory(db)
}
