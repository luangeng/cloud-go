package handler

import (
	k8s "cloud/vender/k8s"

	"github.com/gin-gonic/gin"
)

func ListNs(c *gin.Context) {
	list := k8s.ListNs()
	c.JSON(200, list)
}

func CreateNs(c *gin.Context) {
	k8s.CreateNs("name")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
