package router

import (
	"github.com/Ricardo-Cardozo/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializaHandler()

	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)

		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
