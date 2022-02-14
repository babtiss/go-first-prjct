package repository

import (
	"github.com/jmoiron/sqlx"
	todo "go-application/model"
)

const (
	usersTableName      = "users"
	todoListsTableName  = "todo_lists"
	usersListsTableName = "users_lists"
	todoItemsTableName  = "todo_items"
	listsItemsTableName = "lists_items"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(listId int, item todo.Item) (int, error)
	GetAll(userId, listId int) ([]todo.Item, error)
	GetById(userId, itemId int) (todo.Item, error)
	Delete(userId, listId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
