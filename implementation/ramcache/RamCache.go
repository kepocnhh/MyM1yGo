package ramcache

import (
	//"fmt"
	"github.com/kepocnhh/MyM1yGo/business/cores/transactions"
    "github.com/kepocnhh/MyM1yGo/business/dao"
    "github.com/kepocnhh/MyM1yGo/business/dao/models"
)

type ramCache struct {
	transactionsData models.Transactions
}

func (r *ramCache) GetTransactions() models.Transactions {
	return r.transactionsData
}

func New() dao.DAO {
	return &ramCache{&transactionsModel{make([]transactions.Model, 0)}}
}

type transactionsModel struct {
	data []transactions.Model
}

func (t *transactionsModel) GetAll() []transactions.Model {
	return t.data
}
func (t *transactionsModel) Add(model transactions.Model) {
	t.data = append(t.data, model)
}

