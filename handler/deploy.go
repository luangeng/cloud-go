package handler

import (
	model "cloud/model"
	deploy "cloud/vender/k8s"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDeploy(c *gin.Context) {
	var param model.Deploy1
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	deploy.CreateDeploy(param)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func ListDeploy(c *gin.Context) {
	list, _ := deploy.ListDeploy("default")
	c.JSON(200, list)
}

func ListDeployDetail(c *gin.Context) {
	list, _ := deploy.ListDeploy("default")
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
