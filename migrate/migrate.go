package migrate

import (
	"ggfighter-server/models"

	"github.com/jinzhu/gorm"
)

// Exec : migrate execute
func Exec(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)

	// db.Model(&models.Tournament{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
