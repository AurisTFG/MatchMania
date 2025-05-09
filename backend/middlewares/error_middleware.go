package middlewares

import (
	"fmt"
	"runtime/debug"

	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				handlePanic(c, err)
			}
		}()

		c.Next()
	}
}

func handlePanic(c *gin.Context, err any) {
	stack := debug.Stack()
	fmt.Printf("[ERROR] Panic recovered: %v\n%s\n", err, string(stack))

	if gin.Mode() == gin.DebugMode {
		errMsg := formatErrorMessage(err)
		r.InternalServerError(c, "Internal Server Error: "+errMsg)
	} else {
		r.InternalServerError(c, "Internal Server Error. Please try again later.")
	}

	c.Abort()
}

func formatErrorMessage(err any) string {
	if e, ok := err.(error); ok {
		return e.Error()
	}
	return fmt.Sprintf("%v", err)
}
