package repository

import (
	"log"

	"trungem.com/hoc-gin/internal/models"
)

type SQLUserRepository struct {
}

func NewSQLUserRepository() UserRepository {
	return &SQLUserRepository{}
}

func (ur *SQLUserRepository) Create(user models.User) {
	log.Println("Create User")
}
func (ur *SQLUserRepository) FindById(id int) {
	log.Println("Find User")
}
