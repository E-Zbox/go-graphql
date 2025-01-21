package models

import (
	"graphql-tutorial/src/graph/model"
	"graphql-tutorial/src/models"
	"graphql-tutorial/src/utils/models/interfaces"
	"strconv"
	"time"
)

func CreateUser(payload interfaces.CreateUserInput) model.UserResponse {
	response := model.UserResponse{
		Error:   "",
		Success: false,
	}

	var user = models.User{
		Age:       int(payload.Age),
		Country:   payload.Country,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Role:      payload.Role,
	}

	if err := models.DBInstance.Create(&user).Error; err != nil {
		response.Error = err.Error()
	} else {
		response.Data = &model.User{
			ID:        strconv.Itoa(int(user.ID)),
			Age:       int32(user.Age),
			Country:   user.Country,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format(time.UnixDate),
			UpdatedAt: user.UpdatedAt.Format(time.UnixDate),
		}
		response.Success = true
	}

	return response
}

func FindUser(payload interfaces.GetUserInput) model.UserResponse {
	response := model.UserResponse{
		Error:   "",
		Success: false,
	}

	var user = models.User{}

	switch {
	case len(payload.Email) > 0:
		if err := models.DBInstance.Where("email=", payload.Email).First(&user).Error; err != nil {
			response.Error = err.Error()
		} else {
			response.Data = &model.User{
				ID:        strconv.Itoa(int(user.ID)),
				Age:       int32(user.Age),
				Country:   user.Country,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Role:      user.Role,
				CreatedAt: user.CreatedAt.Format(time.UnixDate),
				UpdatedAt: user.UpdatedAt.Format(time.UnixDate),
			}
			response.Success = true
		}
	case payload.ID > 0:
		if err := models.DBInstance.Where("id=?", payload.ID).First(&user).Error; err != nil {
			response.Error = err.Error()
		} else {
			response.Data = &model.User{
				ID:        strconv.Itoa(int(user.ID)),
				Age:       int32(user.Age),
				Country:   user.Country,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Role:      user.Role,
				CreatedAt: user.CreatedAt.Format(time.UnixDate),
				UpdatedAt: user.UpdatedAt.Format(time.UnixDate),
			}
			response.Success = true
		}
	default:
		response.Error = "Missing one of `id` or `email` payload to retrieve user"
	}

	return response
}
