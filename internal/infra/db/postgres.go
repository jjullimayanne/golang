package db

import (
	dbInterface "api/internal/core/database/interface"
	db_config "api/internal/core/database/struct"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	connection *sql.DB
	config     *db_config.DBConfig
}

func NewDatabase(config *db_config.DBConfig) *Database {
	return &Database{config: config}
}

func (d *Database) Connect() (*sql.DB, error) {
	dsn := d.config.DSN()
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	d.connection = db
	return db, nil
}

func (d *Database) Close() error {
	if d.connection != nil {
		return d.connection.Close()
	}
	return nil
}

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.connection.QueryRow(query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.connection.Exec(query, args...)
}

var _ dbInterface.Database = (*Database)(nil)
