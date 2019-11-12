package admin

import (
	"api"
	"github.com/gin-gonic/gin"
)

func RegistryUserRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/currentUser", api.GetCurrentUser)
}
