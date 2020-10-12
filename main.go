package main

import (
	"testGin/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title DataShare Api
// @version 1.0
// @description 铭数数据共享平台v1.0 Api
// @contact.name 铭数区块链组
// @contact.email liweihang@mingbyte.com

// @host localhost:10086
// @BasePath /datashare

func main() {
	r := gin.Default()
	c := controller.Controller{}

	share := r.Group("/datashare")
	{
		crypto := share.Group("/querySql")
		{
			crypto.POST("", c.SearchResults)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(":10086")
}
