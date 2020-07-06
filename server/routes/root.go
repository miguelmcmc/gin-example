package routes

import (
	"github.com/gin-gonic/gin"
)

//Root funcion of server
func Root(c *gin.Context) {
	c.JSON(200, "Server working corectly")
	return

}
