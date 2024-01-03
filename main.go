/*
Ethiopian Calendar tool for Golang
Copyright (c) 2022 Yinebeb Tariku <yintar5@gmail.com>
*/
package main

import (
	"github.com/Yinebeb-01/ethiopiandateconverter/api"
	docs "github.com/Yinebeb-01/ethiopiandateconverter/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger API
// @version         1.0
// @description     This is a sample swagger server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	// set swagger info
	//docs.SwaggerInfo.Title = "Swagger Example API"
	//docs.SwaggerInfo.Description = "This is a ethioGrego server."
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "swagger.io"
	//docs.SwaggerInfo.BasePath = "api/v1"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		v1.GET("/", api.HomePage)
		v1.GET("/ad-from-et/:date", api.GetAdFromEt)
		v1.GET("/et-from-ad/:date", api.GetEtFromAd)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
