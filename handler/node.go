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
	z, _ := node.ListNode()
	c.JSON(200, z)
}
