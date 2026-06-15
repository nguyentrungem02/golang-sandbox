package repository

import "trungem.com/hoc-gin/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	FindById(id int, user *models.User) error
}
