package handler

import (
	. "cloud/model"
	pv "cloud/vender/k8s"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPv(c *gin.Context) {
	pvcs, err := pv.ListPv("default")
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
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
	c.JSON(200, Ok(results))
}

func ListPvDetail(c *gin.Context) {
	list, err := pv.ListPv("default")
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(list))
}

func CreatePvc(c *gin.Context) {
	var param Pvc
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	err = pv.CreatePvc(param)
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(nil))
}

func DeletePvc(c *gin.Context) {
	var param Pvc
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	err = pv.DeletePvc(param)
	c.JSON(200, Ok(err))
}
