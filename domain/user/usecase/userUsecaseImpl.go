package usecase

import (
	"errors"
	"test-mekar-backend/domain/user/repository"
	"test-mekar-backend/model"
	"test-mekar-backend/utils"
)

type UserUsecase struct {
	userRepo repository.IUserRepository
}

func NewUserUsecase(userRepo repository.IUserRepository) IUserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u UserUsecase) GetJobs() ([]*model.Job, error) {
	listJob, err := u.userRepo.GetJobs()
	if err != nil {
		return nil, err
	}
	return listJob, nil
}
func (u UserUsecase) GetEducations() ([]*model.Education, error) {
	listEducation, err := u.userRepo.GetEducations()
	if err != nil {
		return nil, err
	}
	return listEducation, nil
}
func (u UserUsecase) GetUsers(keyword, page, limit string) ([]*model.User, error) {
	listUser, err := u.userRepo.GetUsers(keyword, page, limit)
	if err != nil {
		return nil, err
	}
	return listUser, nil
}

func (u UserUsecase) GetUserByID(id string) (*model.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUsecase) CreateNewUser(user *model.User) error {

	err := utils.CheckEmpty(user.UserNIK, user.UserName, user.UserBirth, user.UserJob.JobID, user.UserEducation.EducationID)
	if err != nil {
		return errors.New("Some field is Empty")
	}

	error := u.userRepo.CreateNewUser(user)
	if error != nil {
		return errors.New("Some field is Empty")
	}
	return nil
}

func (u UserUsecase) UpdateUser(id string, user *model.User) error {
	err := utils.CheckEmpty(user.UserNIK, user.UserName, user.UserBirth, user.UserJob.JobID, user.UserEducation.EducationID)
	if err != nil {
		return errors.New("Some field is Empty")
	}
	err = u.userRepo.UpdateUser(id, user)
	if err != nil {
		return errors.New("Can't Update User")
	}
	return nil
}

func (u UserUsecase) DeleteUser(id string) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
