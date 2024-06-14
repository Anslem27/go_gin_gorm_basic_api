package main

import (
	"go_crud/initializers"
	"go_crud/models"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectDb()
}

/* refer to https://gorm.io/docs/index.html
go run migrate/migrate.go
*/

func main() {
	// Migrate the posts schema
	initializers.DB.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})
}
