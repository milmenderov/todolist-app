package repository

import (
	"github.com/jmoiron/sqlx"
	todolist_app "todolist-app"
)

type Authorization interface {
	CreateUser(user todolist_app.User) (int, error)
	GetUser(username, password string) (todolist_app.User, error)
}

type TodoList interface {
	Create(userId int, list todolist_app.TodoList) (int, error)
	GetAll(userId int) ([]todolist_app.TodoList, error)
	GetById(userId, listId int) (todolist_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todolist_app.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todolist_app.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todolist_app.TodoItem, error)
	GetById(userId, itemId int) (todolist_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todolist_app.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
