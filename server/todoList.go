package server

import (
	"sample-todo-app/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func ShowTodoList(dbx *sqlx.DB) gin.HandlerFunc {
	todolistRepository := repository.NewTodoList(dbx)
	return func(c *gin.Context) {
		todoLists, err := todolistRepository.GetAllTodoList(c)

		if err != nil {
			c.Status(500)
			return
		}

		c.HTML(200, "top.html", gin.H{
			"title":     "TODOリスト",
			"todoLists": todoLists,
		})
	}
}
