package main

import (
	"log"

	"github.com/gin-gonic/gin"

	docs "users/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"users/controllers"
	"users/dao"
	"users/services"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8001
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	userController := controllers.UsersController{
		UsersService: services.NewUsersService(),
	}
	dao.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Helloworld)
		user := v1.Group("/location")
		{
			user.GET("", userController.GetUsers)
			user.POST("/user", userController.CreateUsers)
			user.PATCH("/user", userController.UpdateUsers)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}

}
