package server

import (
	"github.com/gin-gonic/gin"
	"ironchip.net/gin-example/server/routes"
)

func any(router *gin.RouterGroup) {
	router.GET("/", routes.Root)
}

func test(router *gin.RouterGroup) {
	router.GET("/random", routes.Random)
	router.POST("/random", routes.RandomByLenght)
}
