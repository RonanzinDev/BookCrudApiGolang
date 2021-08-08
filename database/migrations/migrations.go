package migrations

import (
	"github.com/ronanzindev/book-api-golang/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
