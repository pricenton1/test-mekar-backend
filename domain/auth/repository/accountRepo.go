package accountRepository

import "test-mekar-backend/model"

type IAuthRepo interface {
	ReadAccountByEmail(email string) (*model.Account, error)
	CreateAccount(account *model.Account) error
}
