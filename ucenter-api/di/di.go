package di

import "github.com/gin-gonic/gin"

func init() {
	engine := gin.New()
	authorized := engine.Group("/")
	authorized.POST("/api/v1/user/login", KeyWords())
}
