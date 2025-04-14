package middlewares

import (
	"fmt"
	"runtime/debug"

	r "MatchManiaAPI/responses"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				fmt.Printf("[ERROR] Panic recovered: %v\n%s\n", err, string(stack))

				if gin.Mode() == gin.DebugMode {
					r.InternalServerError(c, "Internal Server Error: "+err.(string))
				} else {
					r.InternalServerError(c, "Internal Server Error")
				}

				c.Abort()
			}
		}()

		c.Next()
	}
}
