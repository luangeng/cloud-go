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
	var list = deploy.ListDeploy("default")
	c.JSON(200, list)
}

func ListDeployDetail(c *gin.Context) {
	var list = deploy.ListDeploy("default")
	c.JSON(200, list)
}

func DeleteDeploy(c *gin.Context) {
	ns := c.Query("namespace")
	name := c.Query("name")
	deploy.DeleteDeploy(ns, name)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
