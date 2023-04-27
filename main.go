package main

import (
	"fmt"
	"os"

	_ "github.com/IngoS11/gorestapp/docs"
	"github.com/IngoS11/gorestapp/model"
	"github.com/gin-gonic/gin"

	"github.com/IngoS11/gorestapp/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Main function for the server
//
//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server.
//	@contact.name	Ingo Sauerzapf
//	@contact.url	https://linked.in/in/ingosauerzapf
//	@contact.email	ingo.sauerzapf@gmail.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/api/v1
func main() {

	// Connect to the database
	// capture connection properties
	// DATABASE_URL environment variable must be set before you are able to use
	// the program
	// var databaseUrl = "postgres://postgres:abcd1234@localhost:5432/postgres"
	var databaseUrl = os.Getenv("DATABASE_URL")
	err := model.ConnectToDB(databaseUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to connect to database::", err)
		os.Exit(1)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		albums := v1.Group("/albums")
		{
			albums.GET("", api.GetAlbum)
			albums.POST("", api.PostAlbum)
			albums.GET(":id", api.GetAlbumById)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
