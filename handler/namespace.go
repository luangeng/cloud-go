package handler

import (
	k8s "cloud/vender/k8s"

	"github.com/gin-gonic/gin"
)

func ListNs(c *gin.Context) {
	list, err := k8s.ListNs()
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(list))
}

func CreateNs(c *gin.Context) {
	err := k8s.CreateNs("name")
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(nil))
}
