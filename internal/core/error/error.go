package coreError

import (
    "errors"
    "fmt"
    "log"
)

var (
    ErrNotFound      = errors.New("recurso não encontrado")
    ErrUnauthorized  = errors.New("não autorizado")
    ErrInvalidInput  = errors.New("entrada inválida")
)

func WrapError(err error, message string) error {
    if err != nil {
        return fmt.Errorf("%s: %w", message, err)
    }
    return nil
}

func Is(err, target error) bool {
    return errors.Is(err, target)
}

func As[T any](err error) (T, bool) {
    var target T
    ok := errors.As(err, &target)
    return target, ok
}

func LogError(err error) {
    if err != nil {
        log.Printf("Erro: %v", err)
    }
}

func NewErrorf(format string, a ...interface{}) error {
    return fmt.Errorf(format, a...)
}

type NotFoundError struct {
    Resource string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s não encontrado", e.Resource)
}

func NewNotFoundError(resource string) error {
    return &NotFoundError{Resource: resource}
}
