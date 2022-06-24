package controllers

import (
	"strconv"

	"github.com/Geekinn/go-micro/app/forms"
	"github.com/Geekinn/go-micro/app/models"
	"github.com/go-ozzo/ozzo-validation/v4"

	"net/http"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type ArticleController struct{}

var articleModel = new(models.ArticleModel)

//Create ...
func (ctrl ArticleController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateArticleForm
	if err := c.ShouldBind(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "json could not be parsed", "errors": err.Error()})
        return
    }

	validationErr := validation.ValidateStruct(&form,
		validation.Field(&form.Title, validation.Required, validation.Length(3, 100)),
		validation.Field(&form.Content, validation.Required, validation.Length(3, 100)),
	)
	if validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message":"Article could not be created", "errors": validationErr})
		return	
	}

	id, err := articleModel.Create(userID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created", "errors": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created", "id": id})
}

//All ...
func (ctrl ArticleController) All(c *gin.Context) {
	userID := getUserID(c)

	results, err := articleModel.All(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

//One ...
func (ctrl ArticleController) One(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}

	data, err := articleModel.One(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

//Update ...
func (ctrl ArticleController) Update(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}

	var form forms.UpdateArticleForm

	if err := c.ShouldBind(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message":"Article could not be updated", "errors":  err.Error()})
        return
    }

	validationErr := validation.ValidateStruct(&form,
		validation.Field(&form.Title, validation.When(form.Content == "", validation.Required.Error("Either Content or Title is required.")), validation.Length(3, 100)),
		validation.Field(&form.Content, validation.When(form.Title == "", validation.Required.Error("Either Content or Title is required.")), validation.Length(3, 100)),
		
	)
	if validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message":"Article could not be updated", "errors": validationErr})
		return	
	}

	err = articleModel.Update(userID, getID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be updated", "errors": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article updated"})
}

//Delete ...
func (ctrl ArticleController) Delete(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}

	err = articleModel.Delete(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})

}
