package handler

import (
	deploy "cloud/vender/k8s"

	"github.com/gin-gonic/gin"
)

func ListDeploy(c *gin.Context) {
	var z = deploy.ListDeploy("default")
	c.JSON(200, z)
}

func ListDeployDetail(c *gin.Context) {
	var z = deploy.ListDeploy("default")
	c.JSON(200, z)
}
