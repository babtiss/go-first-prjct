package service

import (
	todo "go-application/model"
	"go-application/pkg/repository"
)

type TodoListService struct {
	Repo repository.TodoList
}

func NewTodoListService(Repo repository.TodoList) *TodoListService {
	return &TodoListService{Repo: Repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.Repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.Repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.Repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.Repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todo.ListInput) error {
	return s.Repo.Update(userId, listId, input)
}
