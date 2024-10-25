package repositories

import (
    "api/internal/modules/signup/domain/entities"
    "api/internal/modules/signup/domain/repositories"
    "api/internal/core/auth" // Importa a interface Authenticator
    "api/internal/core/error"
)

type UserRepository struct {
    Authenticator auth.Authenticator 
}

func NewUserRepository(authenticator auth.Authenticator) repositories.SignUpRepository {
    return &UserRepository{Authenticator: authenticator}
}

func (repo *UserRepository) CreateUser(user *entities.User) error {
    err := repo.Authenticator.CreateUser(user.Username, user.Email, user.Password)
    if err != nil {
        return coreError.WrapError(err, "falha ao criar usuário no Keycloak")
    }
    return nil
}
