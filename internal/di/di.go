package di

import (
    "api/internal/modules/signup/controllers"
    "api/internal/modules/signup/usecase"   
    "api/internal/modules/signup/data/repositories"
    "api/internal/core"
)

func InjectDependencies(database core.Database) (*controllers.AuthController, error) {
    userRepository := repositories.NewUserRepository(database)

    registerUserUseCase := usecases.NewRegisterUserUseCase(userRepository)

    authController := controllers.NewAuthController(registerUserUseCase)

    return authController, nil
}
