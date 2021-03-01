package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"test-mekar-backend/model"
	"test-mekar-backend/utils"

	guuid "github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepo{db: db}
}

// Getusers -> this function fetch accunt info from database and returns found  users.
func (u *UserRepo) GetUsers(keyword, page, limit string) ([]*model.User, error) {
	queryInput := fmt.Sprintf(utils.SELECT_ALL_USER, page, limit)
	rows, err := u.db.Query(queryInput, "%"+keyword+"%")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	listUser := []*model.User{}
	for rows.Next() {
		p := model.User{}
		err := rows.Scan(&p.UserID, &p.UserNIK, &p.UserName, &p.UserBirth, &p.UserJob.JobID, &p.UserJob.JobLabel, &p.UserEducation.EducationID, &p.UserEducation.EducationLabel, &p.UserStatus, &p.UserCreate)
		if err != nil {
			return nil, err
		}
		listUser = append(listUser, &p)
	}

	return listUser, nil
}

func (u *UserRepo) GetJobs() ([]*model.Job, error) {
	rows, err := u.db.Query(utils.SELECT_ALL_JOB)
	if err != nil {
		return nil, err
	}
	listJob := []*model.Job{}
	for rows.Next() {
		p := model.Job{}
		err := rows.Scan(&p.JobID, &p.JobLabel)
		if err != nil {
			return nil, err
		}
		listJob = append(listJob, &p)
	}
	return listJob, nil
}

func (u *UserRepo) GetEducations() ([]*model.Education, error) {
	rows, err := u.db.Query(utils.SELECT_ALL_EDUCATION)
	if err != nil {
		return nil, err
	}
	listEducation := []*model.Education{}
	for rows.Next() {
		p := model.Education{}
		err := rows.Scan(&p.EducationID, &p.EducationLabel)
		if err != nil {
			return nil, err
		}
		listEducation = append(listEducation, &p)
	}

	return listEducation, nil
}

func (u *UserRepo) GetUserByID(id string) (*model.User, error) {
	stmt, err := u.db.Prepare(utils.SELECT_USER_BY_ID)
	user := model.User{}
	if err != nil {
		return nil, err
	}
	errQuery := stmt.QueryRow(id).Scan(&user.UserID, &user.UserNIK, &user.UserName, &user.UserBirth, &user.UserJob.JobID, &user.UserJob.JobLabel, &user.UserEducation.EducationID, &user.UserEducation.EducationLabel, &user.UserStatus, &user.UserCreate)
	if errQuery != nil {
		return &user, err
	}

	defer stmt.Close()
	return &user, nil
}

func (u *UserRepo) UpdateUser(id string, user *model.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(user.UserNIK, user.UserName, user.UserBirth, user.UserJob.JobID, user.UserEducation.EducationID, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (u *UserRepo) CreateNewUser(user *model.User) error {
	userID := guuid.New()
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(userID, user.UserNIK, user.UserName, user.UserBirth, user.UserJob.JobID, user.UserEducation.EducationID); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *UserRepo) DeleteUser(id string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, user id tidak di temukan")
	}
	tx.Commit()
	return nil
}
