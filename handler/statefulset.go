package handler

import (
	k8s "cloud/vender/k8s"

	"github.com/gin-gonic/gin"
)

func ListStateful(c *gin.Context) {
	list, err := k8s.ListStateful("default")
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(list))
}

func CreateStateful(c *gin.Context) {
	err := k8s.CreateStateful()
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(nil))
}

func DeleteStateful(c *gin.Context) {
	err := k8s.DeleteStateful("default", "demo")
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(nil))
}
