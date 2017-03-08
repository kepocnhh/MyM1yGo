package modules

import "github.com/kepocnhh/MyM1yGo/business/contracts/MainContract"

type mainPresenter struct {
    view MainContract.View
    model MainContract.Model
}

func (p *mainPresenter) NewTransaction(count int) {
    p.model.Add(count)
    p.view.UpdateBalance(p.model.GetBalance())
}

func (p *mainPresenter) Update() {
    p.view.UpdateList(p.model.GetAll())
}

func NewMainPresenter(v MainContract.View, m MainContract.Model) MainContract.Presenter {
	return &mainPresenter{v, m}
}