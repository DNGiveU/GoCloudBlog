package controller

import (
	"api"
	"constants"
	"github.com/gin-gonic/gin"
)

func registryLoginRoute(engine *gin.Engine) {
	engine.POST(constants.ADMIN_API + "/login/account", api.LoginAccount)
}
