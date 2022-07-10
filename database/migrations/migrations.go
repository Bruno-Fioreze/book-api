package migrations

import (
	"github.com/bruno-fioreze/book_api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
