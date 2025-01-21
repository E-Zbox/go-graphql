package interfaces

import "graphql-tutorial/src/graph/model"

type CreateUserInput struct {
	Age       int32  `json:"age"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

type GetUserInput struct {
	ID    uint
	Email string
}

type UserResponse struct {
	Data    *model.User
	Error   string
	Success bool
}
