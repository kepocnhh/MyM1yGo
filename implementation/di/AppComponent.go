package di

import (
	//"fmt"
	"github.com/kepocnhh/MyM1yGo/business/dao"
)

type AppComponent interface {
	GetDataAccess() dao.DAO
}
