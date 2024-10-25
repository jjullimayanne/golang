package repositories

import (
    "api/internal/modules/signup/domain/entities"
    "api/internal/modules/signup/domain/repositories"  
    "api/internal/core/database"
    "api/internal/core/error"
)

//avoid sql injection

const (
    createUserQuery = `INSERT INTO users (username, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)`
    getUserByUsernameQuery = `SELECT id, username, first_name, last_name, email, password FROM users WHERE username = $1`
)

type UserRepository struct {
    db core.Database
}

func NewUserRepository(db core.Database) repositories.SignUpRepository {  
    return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *entities.User) error {
    _, err := repo.db.Exec(createUserQuery, user.Username, user.FirstName, user.LastName, user.Email, user.Password)
    if err != nil {
        return coreError.WrapError(err, "erro ao criar o usuário")
    }
    return nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*entities.User, error) {
    row := repo.db.QueryRow(getUserByUsernameQuery, username)

    var user entities.User
    err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password)
    if err != nil {
        return nil, coreError.WrapError(err, "erro ao buscar o usuário")
    }

    return &user, nil
}
