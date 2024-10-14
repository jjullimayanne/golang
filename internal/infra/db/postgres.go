package db

import (
    "database/sql"
    "api/configs/db" 
    _ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresDB struct {
    connection *sql.DB
    config     *configs.DBConfig 
}

func NewPostgresDB(config *configs.DBConfig) *PostgresDB {
    return &PostgresDB{config: config}
}

func (p *PostgresDB) Connect() (*sql.DB, error) {
    dsn := p.config.DSN() 
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    p.connection = db
    return db, nil
}

func (p *PostgresDB) Close() error {
    if p.connection != nil {
        return p.connection.Close()
    }
    return nil
}

func (p *PostgresDB) QueryRow(query string, args ...interface{}) *sql.Row {
    return p.connection.QueryRow(query, args...)
}

func (p *PostgresDB) Exec(query string, args ...interface{}) (sql.Result, error) {
    return p.connection.Exec(query, args...)
}
