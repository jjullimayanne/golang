package entities

import "time"

type TransactionType string

const (
    Earned TransactionType = "earned"
    Spent  TransactionType = "spent"
)

type Transaction struct {
    ID          int
    UserID      string
    Type        TransactionType
    Amount      float64
    Description string
    Date        time.Time
}
