package modules

import (
	"time"
	"math/rand"
	"github.com/kepocnhh/MyM1yGo/business/contracts/MainContract"
    "github.com/kepocnhh/MyM1yGo/business/cores/transactions"
    "github.com/kepocnhh/MyM1yGo/business/dao/models"
)

type mainModel struct {
	transactionsData models.Transactions
}

func (m *mainModel) Add(count int) {
	m.transactionsData.Add(NewTransaction(rand.Int(), time.Now().UnixNano() / int64(time.Millisecond), count))
}
func (m *mainModel) GetBalance() int {
	b := 0
	for _, v := range m.transactionsData.GetAll() {
		b += v.GetCount()
	}
    return b
}
func (m *mainModel) GetAll() []transactions.Model {
    return m.transactionsData.GetAll()
}

func NewMainModel(t models.Transactions) MainContract.Model {
	return &mainModel{t}
}