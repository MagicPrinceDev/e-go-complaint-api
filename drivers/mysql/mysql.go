package mysql

import (
	"e-complaint-api/drivers/indonesia_area_api/regency"
	"e-complaint-api/drivers/mysql/seeder"
	"e-complaint-api/entities"
	"fmt"
	"time"

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
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	// Creating the DSN with loc parameter
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		loc.String(),
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
	db.AutoMigrate(entities.NewsLike{})
	db.AutoMigrate(entities.NewsComment{})
	db.AutoMigrate(entities.ComplaintActivity{})
	db.AutoMigrate(entities.Faq{})
	db.AutoMigrate(entities.Chatbot{})
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
	seeder.SeedComplaintLike(db)
	seeder.SeedComplaintActivity(db)
	seeder.SeedFaq(db)
	seeder.SeedNewsComment(db)
	seeder.SeedNewsLike(db)
}
