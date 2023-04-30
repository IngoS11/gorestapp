package main

import (
	"github.com/IngoS11/gorestapp/controllers"
	_ "github.com/IngoS11/gorestapp/docs"
	"github.com/IngoS11/gorestapp/initializers"
	"github.com/IngoS11/gorestapp/middleware"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Initialize the Server with environment variables
func init() {

	// load environment variables
	initializers.LoadEnvVariables()

	// connect to the database
	initializers.ConnectToDb()

	// syncronize database
	initializers.SyncDatabase()

}

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
//	@host			localhost:3000
//	@BasePath		/api/v1
func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", controllers.Signup)
			users.POST("/login", controllers.Login)
			users.GET("/validate", middleware.RequireAuth, controllers.Validate)
		}

		albums := v1.Group("/albums")
		{
			albums.GET("", controllers.GetAllAlbums)
			albums.POST("", controllers.AddAlbum)
			albums.GET(":id", controllers.GetAlbumById)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
