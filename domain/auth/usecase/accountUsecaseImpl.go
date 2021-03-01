package accountUsecase

import (
	"errors"
	accountRepository "test-mekar-backend/domain/auth/repository"
	"test-mekar-backend/model"
	"test-mekar-backend/utils"
)

type authUsecase struct {
	authRepo accountRepository.IAuthRepo
}

func (a authUsecase) Login(email string) (*model.Account, error) {
	auth, err := a.authRepo.ReadAccountByEmail(email)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a authUsecase) Register(account *model.Account) error {
	err := utils.CheckEmpty(account.Email, account.Password)
	if err != nil {
		return errors.New("Some field is Empty")
	}
	error := a.authRepo.CreateAccount(account)
	if error != nil {
		return errors.New("Some field is Empty")
	}
	return nil
}

func NewAuthUsecase(authRepo accountRepository.IAuthRepo) IAuthUsecase {
	return &authUsecase{authRepo: authRepo}
}
