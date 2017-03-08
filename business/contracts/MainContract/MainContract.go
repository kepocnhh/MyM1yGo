package MainContract

import (
    "github.com/kepocnhh/MyM1yGo/business/cores/transactions"
)

type View interface {
    UpdateBalance(int)
    UpdateList([]transactions.Model)
}
type Presenter interface {
    Update()
    NewTransaction(int)
}
type Model interface {
    GetAll() []transactions.Model
    GetBalance() int
    Add(int)
}