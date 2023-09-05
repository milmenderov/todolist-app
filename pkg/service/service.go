package service

import (
	todolist_app "todolist-app"
	"todolist-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todolist_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todolist_app.TodoList) (int, error)
	GetAll(userId int) ([]todolist_app.TodoList, error)
	GetById(userId, listId int) (todolist_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todolist_app.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todolist_app.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todolist_app.TodoItem, error)
	GetById(userId, itemId int) (todolist_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todolist_app.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
