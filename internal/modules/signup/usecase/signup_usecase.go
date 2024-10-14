package usecases

import (
    "api/internal/modules/signup/domain/entities"
    "api/internal/modules/signup/domain/repositories"
    "errors"
)

type RegisterUserUseCase struct {
    SignUpRepository repositories.SignUpRepository
}

func NewRegisterUserUseCase(signUpRepo repositories.SignUpRepository) *RegisterUserUseCase {
    return &RegisterUserUseCase{
        SignUpRepository: signUpRepo,
    }
}

func (uc *RegisterUserUseCase) Register(user *entities.User) error {
    existingUser, err := uc.SignUpRepository.GetUserByUsername(user.Username)
    if err != nil {
        return err
    }
    if existingUser != nil {
        return errors.New("usuário já existe")
    }

    err = uc.SignUpRepository.CreateUser(user)
    if err != nil {
        return err
    }

    return nil
}
