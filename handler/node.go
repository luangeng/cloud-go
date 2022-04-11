package handler

import (
	node "cloud/vender/k8s"

	. "cloud/model"

	"github.com/gin-gonic/gin"
)

func ListNode(c *gin.Context) {
	list, _ := node.ListNode()
	results := []Node{}
	for _, v := range list.Items {
		node := new(Node)
		node.Name = v.ClusterName
		results = append(results, *node)
	}
	c.JSON(200, results)
}

func ListNodeDetail(c *gin.Context) {
	list, err := node.ListNode()
	if err != nil {
		c.JSON(200, Error(err.Error()))
		return
	}
	c.JSON(200, Ok(list))
}
