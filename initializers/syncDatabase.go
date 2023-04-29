package initializers

import "github.com/IngoS11/gorestapp/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Album{})
	DB.AutoMigrate(&models.User{})
}
