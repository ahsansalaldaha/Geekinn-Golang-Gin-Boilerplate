package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Geekinn/go-micro/app/clients"

	"github.com/gin-gonic/gin"
)

type APIController struct{}

type Todo struct {
    UserId int
	Id int
	Title string
	Completed bool
}

func (ctrl APIController) GetTodo(c *gin.Context) {
	var api clients.API
	api.Init()
	
	todo := Todo{}
	err := api.JSONCall("https://jsonplaceholder.typicode.com/todos/1","GET", &todo)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Print(todo)
	c.JSON(http.StatusOK, todo)
}

func (ctrl APIController) GetGoogle(c *gin.Context)  {
	var api clients.API
	api.Init()
	resp, err := api.Call("https://google.com/","GET")
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	respBytes,_ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	c.JSON(http.StatusOK, string(respBytes))
	
}