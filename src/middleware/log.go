package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

/*
 * 日志
 */
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before
		t := time.Now()

		// invoke
		c.Next()

		// after
		latency := time.Since(t)
		status := c.Writer.Status()
		log.Println("times: " + latency.String() + ", status: " + strconv.Itoa(status))
	}
}
