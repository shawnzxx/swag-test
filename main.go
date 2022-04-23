package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shawnzxx/swag-test/docs"
	"github.com/shawnzxx/swag-test/internal"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo\n"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/todo", internal.GetAllTodos)
	router.GET("/todo/:id", internal.GetTodoByID)
	router.POST("/todo", internal.CreateTodo)
	router.DELETE("/todo", internal.DeleteTodo)

	// run the Gin server
	router.Run()
}
