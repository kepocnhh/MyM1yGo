package dao

import "github.com/kepocnhh/MyM1yGo/business/dao/models"

type DAO interface {
    GetTransactions() models.Transactions
}