package service

import (
	todo "go-application/model"
	"go-application/pkg/repository"
)

type TodoItemService struct {
	itemRepo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{itemRepo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.Item) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.itemRepo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.Item, error) {
	return s.itemRepo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todo.Item, error) {
	return s.itemRepo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.itemRepo.Delete(userId, itemId)
}
