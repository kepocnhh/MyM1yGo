package models

import (
	"github.com/kepocnhh/MyM1yGo/business/cores/transactions"
)

type Transactions interface {
    GetAll() []transactions.Model
    Add(transactions.Model)
}