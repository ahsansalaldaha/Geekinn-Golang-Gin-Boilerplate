package routes

import (
	"github.com/Geekinn/go-micro/app/controllers"
	"github.com/Geekinn/go-micro/app/middlewares"
	"github.com/gin-gonic/gin"
)

func APIRoutes(r *gin.Engine)  {
	v1 := r.Group("/v1")
	{
		/*** START USER Routes ***/
		userController := new(controllers.UserController)

		v1.POST("/user/login", userController.Login)
		v1.POST("/user/register", userController.Register)
		v1.GET("/user/logout", userController.Logout)

		/*** START AUTH Routes ***/
		authController := new(controllers.AuthController)

		//Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", authController.Refresh)

		/*** START Article ***/
		articleController := new(controllers.ArticleController)

		articleRoute := v1.Group("/articles")
		articleRoute.Use(middlewares.TokenAuthMiddleware())
		{
			articleRoute.POST("/",  articleController.Create)
			articleRoute.GET("/all",  articleController.All)
			articleRoute.GET("/",  articleController.Paginate)
			articleRoute.GET("/:id",  articleController.One)
			articleRoute.PUT("/:id",  articleController.Update)
			articleRoute.DELETE("/:id",  articleController.Delete)
		}

		/*** START API Routes ***/
		apiController := new(controllers.APIController)
		v1.GET("/api/todo", apiController.GetTodo)
		v1.GET("/api/google", apiController.GetGoogle)
		
	}
}