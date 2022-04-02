package handler

import (
	k8s "cloud/vender/k8s"

	"github.com/gin-gonic/gin"
)

func ListStateful(c *gin.Context) {
	list, _ := k8s.ListStateful("default")
	c.JSON(200, list)
}

func CreateStateful(c *gin.Context) {
	k8s.CreateStateful()
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteStateful(c *gin.Context) {
	k8s.DeleteStateful("default", "demo")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
