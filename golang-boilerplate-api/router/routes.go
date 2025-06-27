package router

import (
	docs "github.com/LuisMarchio03/golang-boilerplate-api/docs"
	"github.com/LuisMarchio03/golang-boilerplate-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			handler.ShowOpeningHandler(ctx)
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			handler.CreateOpeningHandler(ctx)
		})
		v1.DELETE("/opening", func(ctx *gin.Context) {
			handler.DeleteOpeningHandler(ctx)
		})
		v1.PUT("/opening", func(ctx *gin.Context) {
			handler.UpdateOpeningHandler(ctx)
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			handler.ListOpeningsHandler(ctx)
		})
	}
	// Intialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
	))
}
