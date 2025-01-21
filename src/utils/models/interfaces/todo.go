package interfaces

import (
	"graphql-tutorial/src/graph/model"
)

type CreateTodoInput struct {
	Started bool
	Text    string
	UserId  uint
}

type TodoResponse struct {
	Data    *model.Todo
	Error   string
	Success bool
}

type TodosResponse struct {
	Data    *[]model.Todo
	Error   string
	Success bool
}
