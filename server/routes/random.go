package routes

import (
	"github.com/gin-gonic/gin"
	"ironchip.net/gin-example/models"
	"ironchip.net/gin-example/random"
)

//Random generation
func Random(c *gin.Context) {
	lenght := 10
	randomstring, err := random.GenerateRandomString(lenght)
	if err != nil {
		c.JSON(500, "Error generating random number")
		return
	}

	c.JSON(200, models.Random{
		Random: randomstring,
	})
}
