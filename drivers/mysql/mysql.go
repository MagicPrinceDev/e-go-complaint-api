package mysql

import (
	"e-complaint-api/drivers/indonesia_area_api/regency"
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

	regencyAPI := regency.NewRegencyAPI()

	Migration(db)
	Seeder(db, regencyAPI)

	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(entities.Admin{})
	db.AutoMigrate(entities.User{})
	db.AutoMigrate(entities.Category{})
	db.AutoMigrate(entities.Regency{})
	db.AutoMigrate(entities.Complaint{})
	db.AutoMigrate(entities.ComplaintFile{})
	db.AutoMigrate(entities.ComplaintProcess{})
	db.AutoMigrate(entities.Discussion{})
	db.AutoMigrate(entities.News{})
	db.AutoMigrate(entities.NewsFile{})
	db.AutoMigrate(entities.ComplaintLike{})
}

func Seeder(db *gorm.DB, regencyAPI entities.RegencyIndonesiaAreaAPIInterface) {
	seeder.SeedAdmin(db)
	seeder.SeedUser(db)
	seeder.SeedCategory(db)
	seeder.SeedRegencyFromAPI(db, regencyAPI)
	seeder.SeedComplaint(db)
	seeder.SeedComplaintFile(db)
	seeder.SeedComplaintProcess(db)
	seeder.SeedDiscussion(db)
	seeder.SeedNews(db)
	seeder.SeedNewsFile(db)
}
