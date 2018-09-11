package route

import "github.com/gin-gonic/gin"

// SetRoute set route
func SetRoute(e *gin.Engine) {
	e.Group("/user")
	{
		e.POST("/test")
	}
}
