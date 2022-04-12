package handler

import (
	"cloud/vender/k8s"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	go k8s.TestLock()

	c.JSON(200, gin.H{
		"message": "ok",
	})

}
