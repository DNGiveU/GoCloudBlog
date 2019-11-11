package api

import "net/http"
import "github.com/gin-gonic/gin"
import "model/mind"

/**
 * 获取脑图数据
 */
func GetMindData(c *gin.Context) {
	c.JSON(http.StatusOK, getFakeMindData())
}

func getFakeMindData() mind.Mind {
	rootNode := mind.MindNode{
		Label:    "脑图",
		Children: nil,
	}

	rootNode.Children = []mind.MindNode{mind.MindNode{
		Label:    "卷积神经网络",
		Side:     "right",
		Children: nil,
	}}

	return mind.Mind{Node: []mind.MindNode{rootNode}}
}
