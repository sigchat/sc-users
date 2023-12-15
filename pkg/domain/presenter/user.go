package presenter

import (
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/sigchat/sc-users/pkg/items"
)

type UserPresenter struct {
}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) UserSearchResponse(foundUsers []model.User) (translated []model.User) {
	usersList := items.List[model.User](foundUsers)
	usersList.Filter(func(item model.User, index int) bool {
		var translatedStatus string
		switch item.Status {
		case model.StatusAFK:
			translatedStatus = "Нет на месте"
		case model.StatusOnline:
			translatedStatus = "В сети"
		case model.StatusDND:
			translatedStatus = "Не беспокоить"
		default:
			translatedStatus = "Не в сети"
		}
		usersList[index].Status = model.Status(translatedStatus)
		return false
	})
	return usersList.Slice()
}
