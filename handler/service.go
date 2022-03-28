package handler

import (
	model "cloud/model"
	k8s "cloud/vender/k8s"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListService(c *gin.Context) {
	k8s.ListService("default")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func CreateService(c *gin.Context) {
	var param model.ServiceParam
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	k8s.CreateService(param)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteService(c *gin.Context) {
	k8s.DeleteService("default", "test")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
