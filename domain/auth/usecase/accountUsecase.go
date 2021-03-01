package accountUsecase

import "test-mekar-backend/model"

type IAuthUsecase interface {
	Login(email string) (*model.Account, error)
	Register(account *model.Account) error
}
