package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/Geekinn/go-micro/app/controllers"
)

//TokenAuthMiddleware ...
//JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
var auth = new(controllers.AuthController)
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}