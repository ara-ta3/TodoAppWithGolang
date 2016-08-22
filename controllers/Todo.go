package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Repository repositories.TodoRepository
}

func (t *Todo) ShowAll(c *gin.Context) {
	todos, e := t.Repository.FindAll()
	if e != nil {
		c.String(500, "%s", e)
		return
	}

	c.JSON(200, gin.H{
		"error": nil,
		"todos": todos,
	})
}

func (t *Todo) Show(c *gin.Context) {
	id := c.Param("id")

	tid, e := strconv.Atoi(id)
	if e != nil {
		c.JSON(400, gin.H{
			"error": e,
		})
		return
	}

	todo, e := t.Repository.FindTodo(tid)
	if e != nil {
		c.JSON(500, gin.H{
			"error": e,
		})
		return
	}

	if todo == nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("Todo (id: %s) is not found", id),
		})
		return
	}

	c.JSON(200, gin.H{
		"todo":  todo,
		"error": nil,
	})
}

func (t *Todo) Create(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	if title == "" || description == "" {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("title (%s) and description (%s) cannot be empty", title, description),
		})
		return
	}
	e := t.Repository.PutTodo(&models.Todo{
		Title:       title,
		Description: description,
	})
	if e != nil {
		c.JSON(500, gin.H{"error": e})
		return
	}

	c.JSON(200, gin.H{"error": nil})
}

func (t *Todo) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("Todo (id: %s) is not found", id),
		})
		return
	}
	tid, e := strconv.Atoi(id)
	if e != nil {
		c.JSON(400, gin.H{"error": e})
		return
	}

	todo, e := t.Repository.FindTodo(tid)
	if e != nil {
		c.JSON(500, gin.H{
			"error": e,
		})
		return
	}

	done := c.PostForm("done")
	if done == "" {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("done (%s) cannot be empty", done),
		})
		return
	}

	d, e := strconv.ParseBool(done)
	if e != nil {
		c.JSON(400, gin.H{
			"error": e,
		})
		return
	}

	todo.Done = d
	e = t.Repository.PutTodo(todo)
	if e != nil {
		c.JSON(500, gin.H{"error": e})
		return
	}
	c.JSON(200, gin.H{"error": nil})
}

func (t *Todo) Delete(c *gin.Context) {
	raw := c.PostForm("ids")
	if raw == "" {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Todo (ids: %s) is not found", raw)})
		return
	}
	ids := strings.Split(raw, ",")
	for _, id := range ids {
		tid, e := strconv.Atoi(id)
		if e != nil {
			c.JSON(400, gin.H{"error": e})
			return
		}

		e = t.Repository.RemoveTodo(tid)
		if e != nil {
			c.JSON(500, gin.H{"error": e})
			return
		}
	}
	c.JSON(200, gin.H{"error": nil})
}
