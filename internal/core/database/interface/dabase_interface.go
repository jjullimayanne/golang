package dbInterface

import "database/sql"

type Database interface {
    Connect() (*sql.DB, error)
    Close() error
    QueryRow(query string, args ...interface{}) *sql.Row
    Exec(query string, args ...interface{}) (sql.Result, error)
}
