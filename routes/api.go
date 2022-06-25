package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Geekinn/go-micro/app/controllers"
	"github.com/Geekinn/go-micro/app/middlewares"
)

func APIRoutes(r *gin.Engine)  {
	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		userController := new(controllers.UserController)

		v1.POST("/user/login", userController.Login)
		v1.POST("/user/register", userController.Register)
		v1.GET("/user/logout", userController.Logout)

		/*** START AUTH ***/
		authController := new(controllers.AuthController)

		//Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", authController.Refresh)

		/*** START Article ***/
		articleController := new(controllers.ArticleController)

		articleRoute := v1.Group("/articles")
		articleRoute.Use(middlewares.TokenAuthMiddleware())
		{
			articleRoute.POST("/",  articleController.Create)
			articleRoute.GET("/",  articleController.All)
			articleRoute.GET("/:id",  articleController.One)
			articleRoute.PUT("/:id",  articleController.Update)
			articleRoute.DELETE("/:id",  articleController.Delete)
		}
		
	}
}