package auth

import (
	"net/http"
)

type Authenticator interface {
	ValidateToken(r *http.Request) (bool, error)
	GetUserRoles(r *http.Request) ([]string, error)
	Middleware(next http.Handler) http.Handler
	CreateUser(username, email, password string) error
}
