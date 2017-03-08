package presentation

import (
	"fmt"
    "github.com/kepocnhh/MyM1yGo/business/cores/transactions"
)

type MainView struct {
}

func (v *MainView) UpdateBalance(balance int) {
    fmt.Println("balance", balance)
}
func (v *MainView) UpdateList(list []transactions.Model) {
	if len(list) == 0 {
    	fmt.Println("\ttransactions list is empty!")
    	return
	}
	for i, v := range list {
    	fmt.Println(" ", i, ") id:", v.GetId(), " count:", v.GetCount(), "\tdate:", v.GetDate())
	}
}