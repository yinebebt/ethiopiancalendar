/*
Ethiopian Calendar tool for Golang
Copyright (c) 2022 Yinebeb Tariku <yintar5@gmail.com>
*/
package main

import (
	"github.com/Yinebeb-01/ethiopiandateconverter/config"
	"github.com/Yinebeb-01/ethiopiandateconverter/docs"
	"github.com/Yinebeb-01/ethiopiandateconverter/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           EthioGrego
// @version         1.0
// @description     Ethiopian to Gregorian date converter server.
// @termsOfService  https://st-son.com/terms/
// @contact.name   API Support
// @contact.url    https://st-son.com/support
// @contact.email  support@st-son.com
// @license.name  Apache 2.0
// @license.url   https://www.apache.org/licenses/LICENSE-2.0.html
// @host      calendar.st-son.com
// @BasePath  /v1
// @securityDefinitions.basic  BasicAuth
func main() {
	config.InitConfig("config/config.yaml")
	docs.SwaggerInfo.Schemes = viper.GetStringSlice("server.schemes")
	docs.SwaggerInfo.Host = viper.GetString("server.host")

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"
	router.Static("/assets", "./internal/assets")
	router.StaticFile("favicon.ico", "internal/assets/favicon.ico")
	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		v1.GET("", api.HomePage)
		v1.GET("/et-to-ad/:date", api.Gregorian)
		v1.GET("/ad-to-et/:date", api.Ethiopian)

		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Run(":8080")
}
