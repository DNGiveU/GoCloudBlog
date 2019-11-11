package controller

import (
	"api"
	"github.com/gin-gonic/gin"
)

func registryMindRoute(engine *gin.Engine) {
	engine.GET("/getMindData", api.GetMindData)
}
