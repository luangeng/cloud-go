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
	c.JSON(200, Ok(nil))
}

func ListDeploy(c *gin.Context) {
	list, err := deploy.ListDeploy("default")
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(list))
}

func ListDeployDetail(c *gin.Context) {
	list, err := deploy.ListDeploy("default")
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(list))
}

func DeleteDeploy(c *gin.Context) {
	ns := c.Query("namespace")
	name := c.Query("name")
	err := deploy.DeleteDeploy(ns, name)
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(nil))
}
