package usecase

import "test-mekar-backend/model"

type IUserUsecase interface {
	GetJobs() ([]*model.Job, error)
	GetEducations() ([]*model.Education, error)
	GetUsers(keyword, page, limit string) ([]*model.User, error)
	GetUserByID(id string) (*model.User, error)
	CreateNewUser(user *model.User) error
	UpdateUser(id string, user *model.User) error
	DeleteUser(id string) error
}
