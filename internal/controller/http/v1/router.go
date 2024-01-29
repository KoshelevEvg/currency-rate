package v1

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	api := router.Group("/api/v1")
	{
		api.GET("/currency", GetCur)
	}
	return router
}
