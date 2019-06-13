package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	//_ "./docs" // docs is generated by Swag CLI, you have to import it.
	_ "github.com/NirmalVatsyayan/PersonalProjects/SwaggerTest/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name Nirmal Vatsyayan
// @contact.url http://www.swagger.io/support
// @contact.email nirmal.vatsyayan1990@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := gin.New()

	config := &ginSwagger.Config{
		URL: "http://localhost:3000/swagger/doc.json", //The url pointing to API definition
	}
	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	r.Run(":3000")
}
