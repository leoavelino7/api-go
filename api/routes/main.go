package routes

import (
	"api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/v1")

	tweetController := controllers.NewTweetController()
	cepController := controllers.NewCepController()

	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweets", tweetController.Create)
		v1.DELETE("/tweets/:id", tweetController.Delete)
	}

	{
		v1.GET("/ceps", cepController.FindAll)
		v1.GET("/ceps/:cep", cepController.FindByCep)
	}

	return v1
}