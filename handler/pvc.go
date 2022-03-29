package handler

import (
	. "cloud/model"
	pv "cloud/vender/k8s"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPv(c *gin.Context) {
	var pvcs = pv.ListPv("default")
	results := []Pv{}
	for _, value := range pvcs {
		// fmt.Printf("idx: %d, value: %d\n", idx, value)
		pv := new(Pv)
		pv.Name = value.Spec.VolumeName
		pv.StorageClass = *value.Spec.StorageClassName
		pv.Capacity = value.Status.Capacity.Storage().ToDec().String()
		pv.AccessModes = string(value.Spec.AccessModes[0])
		pv.Status = string(value.Status.Phase)
		results = append(results, *pv)
	}
	c.JSON(200, results)
}

func ListPvDetail(c *gin.Context) {
	var z = pv.ListPv("default")
	c.JSON(200, z)
}

func CreatePvc(c *gin.Context) {
	var param Pvc
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	pv.CreatePvc(param)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeletePvc(c *gin.Context) {
	var param Pvc
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	pv.DeletePvc(param)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}