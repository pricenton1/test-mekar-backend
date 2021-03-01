package controller

import (
	"log"
	"net/http"
	"test-mekar-backend/domain/user/usecase"
	"test-mekar-backend/middleware"
	"test-mekar-backend/model"
	"test-mekar-backend/utils"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func UserController(userUsecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (u *UserHandler) UserAPI(r *mux.Router) {
	users := r.PathPrefix("/users").Subrouter()
	users.Use(middleware.TokenValidationMiddleware)
	users.HandleFunc("", u.GetUsers).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}").Methods(http.MethodGet)
	users.HandleFunc("/{id}", u.GetUserId).Methods(http.MethodGet)
	users.HandleFunc("/register", u.CreateUser).Methods(http.MethodPost)
	users.HandleFunc("/update/{id}", u.UpdateUser).Methods(http.MethodPut)
	users.HandleFunc("/delete/{id}", u.DeleteUser).Methods(http.MethodDelete)

	etcRoute := r.PathPrefix("").Subrouter()
	users.Use(middlewares.TokenValidationMiddleware)
	etcRoute.HandleFunc("/jobs", u.ReadJob).Methods(http.MethodGet)
	etcRoute.HandleFunc("/educations", u.ReadEducation).Methods(http.MethodGet)
}

func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var page string = mux.Vars(r)["page"]
	var limit string = mux.Vars(r)["limit"]
	var keyword string = mux.Vars(r)["keyword"]

	users, err := u.userUsecase.GetUsers(keyword, page, limit)
	if err != nil {
		utils.ResponseWithoutPayload(w, http.StatusBadRequest)
	} else {
		utils.Response(w, http.StatusOK, users)
	}
}

func (u *UserHandler) GetUserId(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	user, err := u.userUsecase.GetUserByID(id)
	if err != nil {
		utils.ResponseWithoutPayload(w, http.StatusBadRequest)
	} else {
		utils.Response(w, http.StatusOK, user)
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := utils.JsonDecoder(&user, r)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	err = u.userUsecase.CreateNewUser(&user)
	log.Print(err)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.ResponseWithoutPayload(w, http.StatusCreated)
}
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	id := utils.DecodePathVariabel("id", r)
	if len(id) == 0 {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	err := utils.JsonDecoder(&user, r)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	err = u.userUsecase.UpdateUser(id, &user)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.ResponseWithoutPayload(w, http.StatusOK)
}
func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) == 0 {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	err := u.userUsecase.DeleteUser(id)
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadGateway)
		return
	}
	utils.ResponseWithoutPayload(w, http.StatusOK)
}

func (u *UserHandler) ReadJob(w http.ResponseWriter, r *http.Request) {
	job, err := u.userUsecase.GetJobs()
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	utils.Response(w, http.StatusOK, job)
}

func (u *UserHandler) ReadEducation(w http.ResponseWriter, r *http.Request) {
	education, err := u.userUsecase.GetEducations()
	if err != nil {
		utils.ResponseWithoutPayload(
			w, http.StatusBadRequest)
		return
	}
	utils.Response(w, http.StatusOK, education)
}
