package route

import "github.com/gin-gonic/gin"

func SetRoute(e *gin.Engine) {
	e.Group("/user")
	{
		e.POST("/test")
	}
}
