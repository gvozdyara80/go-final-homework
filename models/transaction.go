package models

import "time"

type TransactionType string

const (
	Income   TransactionType = "income"
	Expense  TransactionType = "expense"
	Transfer TransactionType = "transfer"
)

type Transaction struct {
	Id          int             `json:"id"`
	Amount      float64         `json:"amount"`
	Currency    string          `json:"currency"`
	Type        TransactionType `json:"type"`
	Category    string          `json:"category"`
	Description string          `json:"description"`
	Date        time.Time       `json:"date"`
	UserId      int             `json:"user_id"`
}

type TransactionRepository interface {
	AddTransaction(transaction Transaction) (int, error)
	GetTransactionById(id int) (Transaction, error)
	GetAllTransactions() ([]Transaction, error)
	UpdateTransaction(id int, transaction Transaction) error
	DeleteTransaction(id int) error
}