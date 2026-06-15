package dto

import "trungem.com/user-manager/internal/models"

type UserDTO struct {
	UUID   string `json:"uuid"`
	Name   string `json:"full_name"`
	Email  string `json:"email_address"`
	Age    int    `json:"age"`
	Status string `json:"status"`
	Level  string `json:"level"`
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"omitempty,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UuidParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

type GetAllUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   int    `form:"page" binding:"omitempty,gte=1"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func (in *CreateUserInput) MapCreateInputToModel() models.User {
	return models.User{
		Name:     in.Name,
		Email:    in.Email,
		Age:      in.Age,
		Password: in.Password,
		Status:   in.Status,
		Level:    in.Level,
	}
}

func (in *UpdateUserInput) MapUpdateInputToModel() models.User {
	return models.User{
		Name:     in.Name,
		Email:    in.Email,
		Age:      in.Age,
		Password: in.Password,
		Status:   in.Status,
		Level:    in.Level,
	}
}

func MapUserToDTO(user models.User) *UserDTO {
	return &UserDTO{
		UUID:   user.UUID,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Status: mapStatus(user.Status),
		Level:  mapLevel(user.Level),
	}
}

func MapUsersToDTO(users []models.User) *[]UserDTO {
	dtos := make([]UserDTO, 0, len(users))

	for _, user := range users {
		dtos = append(dtos, *MapUserToDTO(user))
	}

	return &dtos
}

func mapStatus(status int) string {
	switch status {
	case 1:
		return "Show"
	case 2:
		return "Hide"
	default:
		return "None"
	}
}

func mapLevel(level int) string {
	switch level {
	case 1:
		return "Admin"
	case 2:
		return "Member"
	default:
		return "None"
	}
}
