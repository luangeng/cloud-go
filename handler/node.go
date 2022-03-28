package handler

import (
	node "cloud/vender/k8s"

	. "cloud/model"

	"github.com/gin-gonic/gin"
)

func ListNode(c *gin.Context) {
	var list = node.ListNode()
	results := []Node{}
	for _, v := range list.Items {
		node := new(Node)
		node.Name = v.ClusterName
		results = append(results, *node)
	}
	c.JSON(200, results)
}

func ListNodeDetail(c *gin.Context) {
	var z = node.ListNode()
	c.JSON(200, z)
}
