package controllers

import (
    "encoding/json"
    "net/http"
    "api/internal/modules/signup/domain/entities"
    "api/internal/modules/signup/usecases"
)

type AuthController struct {
    RegisterUserUseCase *usecases.RegisterUserUseCase
}

func NewAuthController(registerUseCase *usecases.RegisterUserUseCase) *AuthController {
    return &AuthController{
        RegisterUserUseCase: registerUseCase,
    }
}

func (controller *AuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
    var userRequest entities.User

    err := json.NewDecoder(r.Body).Decode(&userRequest)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    err = controller.RegisterUserUseCase.Register(&userRequest)
    if err != nil {
        http.Error(w, "Erro ao registrar usuário: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Usuário registrado com sucesso"))
}
