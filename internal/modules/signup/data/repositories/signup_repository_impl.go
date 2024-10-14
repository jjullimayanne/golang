package repositories

import (
    "api/internal/modules/signup/domain/entities"
    "api/internal/modules/signup/domain/repositories"  
    "api/internal/core"
    "fmt"
)

type UserRepository struct {
    db core.Database
}

func NewUserRepository(db core.Database) repositories.SignUpRepository {  
    return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *entities.User) error {
    query := `INSERT INTO users (username, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)`
    _, err := repo.db.Exec(query, user.Username, user.FirstName, user.LastName, user.Email, user.Password)
    if err != nil {
        return fmt.Errorf("erro ao criar o usuário: %v", err)
    }
    return nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*entities.User, error) {
    query := `SELECT id, username, first_name, last_name, email, password FROM users WHERE username = $1`
    row := repo.db.QueryRow(query, username)

    var user entities.User
    err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar o usuário: %v", err)
    }

    return &user, nil
}
