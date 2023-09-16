package repository

import (
	"context"
	"log"
	"sample-todo-app/model"

	"github.com/jmoiron/sqlx"
)

type TodoList interface {
	GetAllTodoList(ctx context.Context) ([]*model.TodoList, error)
}

func NewTodoList(dbx *sqlx.DB) TodoList {
	return &dbTodoList{dbx: dbx}
}

type dbTodoList struct {
	dbx *sqlx.DB
}

func (d *dbTodoList) GetAllTodoList(ctx context.Context) ([]*model.TodoList, error) {
	var todoLists []*model.TodoList

	err := d.dbx.Select(&todoLists, "SELECT * FROM TodoList")
	if err != nil {
		log.Fatal("GetAllTodoListError:", err)
		return nil, err
	}

	for _, tl := range todoLists {
		tl.SetTimes()
	}

	return todoLists, nil
}
