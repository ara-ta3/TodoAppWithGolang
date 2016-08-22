package controllers

import (
	"strconv"

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

	c.JSON(200, gin.H{"todos": todos})
}

func (t *Todo) Show(c *gin.Context) {
	id := c.Param("id")

	tid, e := strconv.Atoi(id)
	if e != nil {
		c.String(400, "%s", e)
		return
	}

	todo, e := t.Repository.FindTodo(tid)
	if e != nil {
		c.String(500, "%s", e)
		return
	}

	if todo == nil {
		c.String(404, "Todo (id: %s) is not found", id)
		return
	}
	str, e := todo.ToJson()
	if e != nil {
		c.String(500, "%s", e)
		return
	}

	c.String(200, "%s", string(str))
}

func (t *Todo) Create(c *gin.Context) {
	title := c.PostForm("title")
	contents := c.PostForm("contents")
	if title == "" || contents == "" {
		c.String(400, "title (%s) and contents (%s) cannot be empty", title, contents)
		return
	}
	e := t.Repository.PutTodo(&models.Todo{
		Title:    title,
		Contents: contents,
	})
	if e != nil {
		c.String(500, "%s", e)
		return
	}

	c.String(200, "%s", string("ok"))
}

func (t *Todo) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(404, "Todo (id: %s) is not found", id)
		return
	}
	tid, e := strconv.Atoi(id)
	if e != nil {
		c.String(400, "%s", e)
		return
	}

	title := c.PostForm("title")
	contents := c.PostForm("contents")
	if title == "" || contents == "" {
		c.String(400, "title (%s) and contents (%s) cannot be empty", title, contents)
		return
	}

	e = t.Repository.PutTodo(&models.Todo{
		ID:       tid,
		Title:    title,
		Contents: contents,
	})
	if e != nil {
		c.String(500, "%s", e)
		return
	}
	c.String(200, "%s", "ok")
}

func (t *Todo) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(404, "Todo (id: %s) is not found", id)
		return
	}
	tid, e := strconv.Atoi(id)
	if e != nil {
		c.String(400, "%s", e)
		return
	}

	e = t.Repository.RemoveTodo(tid)
	if e != nil {
		c.String(500, "%s", e)
		return
	}
	c.String(200, "%s", "ok")
}
