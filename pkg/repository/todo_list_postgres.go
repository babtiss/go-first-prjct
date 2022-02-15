package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "go-application/model"
	"strings"
)

type TodoListPostgres struct {
	DB *sqlx.DB
}

func NewTodoListPostgres(DB *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{DB: DB}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	var id int
	transaction, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}

	createListRequest := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTableName)
	row := transaction.QueryRow(createListRequest, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		err := transaction.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	createUsersListRequest := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTableName)
	_, err = transaction.Exec(createUsersListRequest, userId, id)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, transaction.Commit()

}
func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTableName, usersListsTableName)
	err := r.DB.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.* FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTableName, usersListsTableName)
	err := r.DB.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTableName, usersListsTableName)
	_, err := r.DB.Exec(query, userId, listId)

	return err
}
func (r *TodoListPostgres) Update(userId, listId int, input todo.ListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId += 1
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId += 1
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTableName, setQuery, usersListsTableName, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.DB.Exec(query, args...)
	return err
}
