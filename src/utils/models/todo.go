package models

import (
	"graphql-tutorial/src/graph/model"
	"graphql-tutorial/src/graph/scalar"
	"graphql-tutorial/src/models"
	"graphql-tutorial/src/utils/models/interfaces"
	"strconv"
)

func CreateTodo(payload interfaces.CreateTodoInput) model.TodoResponse {
	todo := models.Todo{
		Completed: false,
		Started:   payload.Started,
		Text:      payload.Text,
		UserId:    payload.UserId,
	}

	models.DBInstance.Create(&todo)

	return model.TodoResponse{
		Data: &model.Todo{
			ID:        strconv.Itoa(int(todo.ID)),
			Completed: todo.Completed,
			Started:   todo.Started,
			Text:      todo.Text,
			UserId:    scalar.Uint(todo.UserId),
			CreatedAt: todo.CreatedAt.Format("2006-01-02T15:04:05 -07:00:00"),
			UpdatedAt: todo.CreatedAt.Format("2006-01-02T15:04:05 -07:00:00"),
		},
		Error:   "",
		Success: true,
	}
}

func FindTodoById(payload interfaces.IDInput) model.TodoResponse {
	response := model.TodoResponse{
		// Data: {},
		Error:   "",
		Success: false,
	}
	var todoExists = models.Todo{}

	if err := models.DBInstance.Where("id=?", payload.ID).First(&todoExists).Error; err != nil {
		response.Error = err.Error()
	} else {
		response.Data = &model.Todo{
			ID:        strconv.Itoa(int(todoExists.ID)),
			Completed: todoExists.Completed,
			Started:   todoExists.Started,
			Text:      todoExists.Text,
			UserId:    scalar.Uint(todoExists.UserId),
			CreatedAt: todoExists.CreatedAt.Format("2006-01-02T15:04:05 -07:00:00"),
			UpdatedAt: todoExists.CreatedAt.Format("2006-01-02T15:04:05 -07:00:00"),
		}
		response.Success = true
	}

	return response
}

func FindTodosByUserId(payload interfaces.IDInput) model.TodosResponse {
	var todos = []models.Todo{}

	response := model.TodosResponse{
		Error:   "",
		Success: false,
	}

	todosArray := []*model.Todo{}

	if err := models.DBInstance.Where("user_id=?", payload.ID).Find(&todos).Error; err != nil {
		response.Error = err.Error()
	} else {
		for _, todo := range todos {
			todosArray = append(todosArray, &model.Todo{
				ID:        strconv.Itoa(int(todo.ID)),
				Completed: todo.Completed,
				Started:   todo.Started,
				Text:      todo.Text,
				UserId:    scalar.Uint(todo.UserId),
				CreatedAt: todo.CreatedAt.Format("2006-01-02T15:04:05 -07:00:00"),
				UpdatedAt: todo.CreatedAt.Format("2006-01-02T15:04:05 -07:00:00"),
			})
		}

		response.Data = todosArray
		response.Success = true
	}

	return response
}

/*


	userId, err := strconv.Atoi(input.UserID)

	if err != nil {
		panic(err)
	}

	response := utilModels.CreateTodo(struct {
		Started bool
		Text    string
		UserId  uint
	}{
		Started: false,
		Text:    input.Text,
		UserId:  uint(userId),
	})

	return response, errors.New(response.Error)
*/
