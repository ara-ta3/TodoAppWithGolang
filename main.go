package main

import (
	"net/http"

	"./controllers"
	"./models"
	"./repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	r := &repositories.TodoRepositoryOnMemory{Data: map[int]*models.Todo{
		1: &models.Todo{
			ID:          1,
			Title:       "ReactでTodoAppを作ってみる",
			Description: "がんばる",
		},
	}}
	t := &controllers.Todo{Repository: r}
	router := gin.Default()
	router.LoadHTMLGlob("views/*.tmpl")
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	api := router.Group("/api")
	{
		api.GET("/todo", t.ShowAll)
		api.GET("/todo/:id", t.Show)
		api.POST("/todo", t.Create)
		api.PUT("/todo/:id", t.Update)
		api.DELETE("/todo", t.Delete)
	}
	router.Run()
}
