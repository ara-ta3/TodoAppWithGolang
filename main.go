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
	router.GET("/todo", t.ShowAll)
	router.GET("/todo/:id", t.Show)
	router.POST("/todo", t.Create)
	router.PUT("/todo/:id", t.Update)
	router.DELETE("/todo/:id", t.Delete)
	router.Run()
}
