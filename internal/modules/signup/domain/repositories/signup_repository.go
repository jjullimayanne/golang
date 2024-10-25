package repositories

import (
	"api/internal/modules/signup/domain/entities"
)

type SignUpRepository interface {
	CreateUser(user *entities.User) error
}
