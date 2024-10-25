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
	err := uc.SignUpRepository.CreateUser(user)
	if err != nil {
		if errors.Is(err, errors.New("usuário já existe")) {
			return errors.New("usuário já existe")
		}
		return err
	}

	return nil
}
