package accountController

import (
	"log"
	"net/http"
	accountUsecase "test-mekar-backend/domain/auth/usecase"
	"test-mekar-backend/model"
	"test-mekar-backend/utils"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	accountUsecase accountUsecase.IAuthUsecase
}

func AccountController(accountUsecase accountUsecase.IAuthUsecase) *AccountHandler {
	return &AccountHandler{accountUsecase: accountUsecase}
}

func (u *AccountHandler) AccountAPI(r *mux.Router) {
	account := r.PathPrefix("/account").Subrouter()
	account.HandleFunc("/register", u.CreateAccount).Methods(http.MethodPost)
	account.HandleFunc("/login", u.Login).Methods(http.MethodPost)

}

func (u *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	err := utils.JsonDecoder(&account, r)
	account.Password = utils.Encrypt([]byte(account.Password))
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	err = u.accountUsecase.Register(&account)
	log.Print(err)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.ResponseWithoutPayload(w, http.StatusCreated)
}

func (u *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	err := utils.JsonDecoder(&account, r)
	if err != nil {
		utils.ResponseWithoutPayload(w, http.StatusBadRequest)
	} else {
		authTemp, err := u.accountUsecase.Login(account.Email)
		if err != nil {
			utils.ResponseWithoutPayload(w, http.StatusBadGateway)
		}
		isValid := utils.CompareEncrypt(authTemp.Password, []byte(account.Password))
		if isValid {
			token, err := utils.JwtEncoder(authTemp.Email, "Rahasia")
			if err != nil {
				utils.ResponseWithoutPayload(w, http.StatusBadRequest)
			}
			authTemp.Token = model.Token{Key: token}
			utils.Response(w, http.StatusOK, authTemp)
		} else {
			utils.ResponseWithoutPayload(w, http.StatusUnauthorized)
		}
	}
}
