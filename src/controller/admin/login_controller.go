package admin

import (
	"api"
	"constants"
	"github.com/gin-gonic/gin"
)

func RegistryLoginRoute(engine *gin.Engine) {
	engine.POST(constants.ADMIN_API + "/login/account", api.LoginAccount)
}
