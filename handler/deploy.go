package handler

import (
	deploy "cloud/vender/k8s"

	"github.com/gin-gonic/gin"
)

func CreateDeploy(c *gin.Context) {
	deploy.CreateDeploy()
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func ListDeploy(c *gin.Context) {
	var z = deploy.ListDeploy("default")
	c.JSON(200, z)
}

func ListDeployDetail(c *gin.Context) {
	var z = deploy.ListDeploy("default")
	c.JSON(200, z)
}

func DeleteDeploy(c *gin.Context) {
	deploy.DeleteDeploy()
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
