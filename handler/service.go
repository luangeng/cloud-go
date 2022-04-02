package handler

import (
	model "cloud/model"
	k8s "cloud/vender/k8s"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListService(c *gin.Context) {
	ns := c.Query("namespace")
	list, err := k8s.ListService(ns)
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(list))
}

func CreateService(c *gin.Context) {
	var param model.ServiceParam
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	err = k8s.CreateService(param)
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(nil))
}

func DeleteService(c *gin.Context) {
	ns := c.Query("namespace")
	name := c.Query("name")
	err := k8s.DeleteService(ns, name)
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(nil))
}
