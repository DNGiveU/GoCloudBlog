package controller

import (
	"api"
	"constants"
	"github.com/gin-gonic/gin"
)

func registryUserRoute(engine *gin.Engine) {
	engine.GET(constants.ADMIN_API + "/currentUser", api.GetCurrentUser)
}
