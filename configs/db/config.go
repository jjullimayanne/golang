package configs

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

type DBConfig struct {
    User     string
    Password string
    Name     string
    Host     string
    Port     string
    SSLMode  string
}

func LoadDBConfig() (*DBConfig, error) {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Nenhum arquivo .env encontrado. Usando variáveis de ambiente.")
    }

    config := &DBConfig{
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        Name:     os.Getenv("DB_NAME"),
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }

    if config.User == "" || config.Password == "" || config.Name == "" || config.Host == "" || config.Port == "" {
        return nil, fmt.Errorf("configurações obrigatórias do banco de dados ausentes")
    }

    return config, nil
}

func (config *DBConfig) DSN() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        config.User, config.Password, config.Host, config.Port, config.Name, config.SSLMode)
}
