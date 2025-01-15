package task

import (
	"fmt"
	"time"
	
	"github.com/gin-gonic/gin"
)

func RouteTasks(router *gin.Engine) {
	t := router.Group("tasks")
	t.Use(CORSMiddleware())
	t.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("\nClient: %s\nTime: %s \n%s [%s]\n%s | Status: %d\n\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			// param.Latency,
			// param.Request.UserAgent(),
			// param.ErrorMessage,
		)
	}))
	
	t.Use(gin.Recovery())
	
	t.GET("")
	t.GET("/:id")
	t.POST("")
	t.PUT("/:id")
	t.DELETE("/:id")
}
