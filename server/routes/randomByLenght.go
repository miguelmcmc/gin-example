package routes

import (
	"github.com/gin-gonic/gin"
	"ironchip.net/gin-example/models"
	"ironchip.net/gin-example/random"
)

//RandomByLenght generation with fixed lenght from body
func RandomByLenght(c *gin.Context) {
	//Get body
	body := models.RandomGenerate{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := body.Validate(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	randomstring, err := random.GenerateRandomString(body.Lenght)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, models.Random{
		Random: randomstring,
	})
}
