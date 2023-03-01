package initializers

import (
	"github.com/LilzBay/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
