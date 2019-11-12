package middleware

import "github.com/gin-gonic/gin"

/*
 * 认证
 * TODO implement
 */
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// invoke
		c.Next()
	}
}
