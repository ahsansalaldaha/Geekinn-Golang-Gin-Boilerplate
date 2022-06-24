package controllers

import (
	"github.com/Geekinn/go-micro/app/forms"
	"github.com/Geekinn/go-micro/app/models"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"net/http"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

//getUserID ...
func getUserID(c *gin.Context) (userID int64) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(int64)
}

//Login ...
func (ctrl UserController) Login(c *gin.Context) {

	var form forms.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "json could not be parsed", "errors": err.Error()})
        return
    }

	validationErr := validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required, validation.Length(3, 100), is.EmailFormat),
		validation.Field(&form.Password, validation.Required, validation.Length(3, 100)),
	)
	if validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message":"Invalid login details"})
		return	
	}


	user, token, err := userModel.Login(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

//Register ...
func (ctrl UserController) Register(c *gin.Context) {

	var form forms.RegisterForm
	if err := c.ShouldBindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "json could not be parsed", "errors": err.Error()})
        return
    }

	validationErr := validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required),
		validation.Field(&form.Email, validation.Required, validation.Length(3, 100),is.EmailFormat),
		validation.Field(&form.Password, validation.Required, validation.Length(3, 100)),
	)
	if validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message":"User could not be created", "errors": validationErr})
		return	
	}

	user, err := userModel.Register(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}

//Logout ...
func (ctrl UserController) Logout(c *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
