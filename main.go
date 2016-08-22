package main

import (
	"./controllers"
	"./models"
	"./repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	r := &repositories.TodoRepositoryOnMemory{Data: map[int]*models.Todo{}}
	t := &controllers.Todo{Repository: r}
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/todo", t.ShowAll)
		api.GET("/todo/:id", t.Show)
		api.POST("/todo", t.Create)
		api.PUT("/todo/:id", t.Update)
		api.DELETE("/todo/:id", t.Delete)
	}
	router.Run()
}
