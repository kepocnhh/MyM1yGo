package modules

import (
	"github.com/kepocnhh/MyM1yGo/business/cores/transactions"
)

type transaction struct {
	id int
	date int64
	count int
}

func (t *transaction) GetId() int {
	return t.id
}
func (t *transaction) GetDate() int64 {
	return t.date
}
func (t *transaction) GetCount() int {
	return t.count
}

func NewTransaction(i int, d int64, c int) transactions.Model {
	return &transaction{i, d, c}
}