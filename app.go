package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
    "strconv"
	"github.com/kepocnhh/MyM1yGo/business/contracts/MainContract"
	"github.com/kepocnhh/MyM1yGo/business/dao"
	"github.com/kepocnhh/MyM1yGo/implementation/modules"
	"github.com/kepocnhh/MyM1yGo/implementation/di"
	"github.com/kepocnhh/MyM1yGo/implementation/ramcache"
	"github.com/kepocnhh/MyM1yGo/presentation"
)

type AppComponent struct {
	Dao dao.DAO
}
func (a AppComponent) GetDataAccess() dao.DAO {
	return a.Dao
}

var appComponent di.AppComponent

var view MainContract.View
var presenter MainContract.Presenter

func main() {
	fmt.Println("\tWelcome!\nJust try add new transaction count...")
	appComponent = AppComponent{ramcache.New()}
	view = &presentation.MainView{}
	presenter = modules.NewMainPresenter(view, modules.NewMainModel(GetAppComponent().GetDataAccess().GetTransactions()))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Compare(s, "q") == 0 {
			return
		}
		if strings.Compare(s, "s") == 0 {
	    	presenter.Update()
	    	continue
		}
	    count, err := strconv.Atoi(s)
	    if err != nil {
	    	fmt.Println("\terror")
	    	continue
	    }
	    presenter.NewTransaction(count)
	}
}

func GetAppComponent() di.AppComponent {
	return appComponent
}
