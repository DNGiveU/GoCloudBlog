package admin

import (
	"api"
	"github.com/gin-gonic/gin"
)

func RegistryLoginRoute(engine *gin.Engine) {
	engine.POST("/login/account", api.LoginAccount)
}
