package config

import (
	"database/sql"
	accountController "test-mekar-backend/domain/auth/controller"
	accountRepository "test-mekar-backend/domain/auth/repository"
	accountUsecase "test-mekar-backend/domain/auth/usecase"
	"test-mekar-backend/domain/user/controller"
	"test-mekar-backend/domain/user/repository"
	"test-mekar-backend/domain/user/usecase"
	"test-mekar-backend/middleware"

	"github.com/gorilla/mux"
)

func NewRoutes(db *sql.DB, r *mux.Router) {
	// USER
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.UserController(userUsecase)
	userController.UserAPI(r)
	// ACCOUNT
	accountRepo := accountRepository.NewAuthRepo(db)
	accountUsecase := accountUsecase.NewAuthUsecase(accountRepo)
	accountController := accountController.AccountController(accountUsecase)
	accountController.AccountAPI(r)

	r.Use(middleware.ActivityLogger)
}
