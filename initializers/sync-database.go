package initializers

import "github.com/bednyak/go-jwt-auth-cookie-v2/models"

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
