package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/domain/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// UserCreateReq represents the data transfer object requested for creating a user.
type UserCreateReq struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	BirthDay  time.Time `json:"birthday"`
	Username  string    `json:"username"`
	Passwd    string    `json:"password"`
}

// UserUpdateReq represents the data transfer object requested for updating a user.
type UserUpdateReq struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birthday"`
	Username  string    `json:"username"`
}

// UserLoginResponse represents the data transfer object response for a logged user.
type UserLoginResponse struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	LastLogin time.Time `json:"last_login"`
}

type UserController struct {
	svc service.UserService
}

func NewUserController(svc service.UserService) UserController {
	return UserController{
		svc: svc,
	}
}

func (uc UserController) create(w http.ResponseWriter, r *http.Request) {
	var dto UserCreateReq
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}
	user := entity.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		BirthDay:  dto.BirthDay,
		Username:  dto.Username,
		Passwd:    dto.Passwd,
	}

	if err := uc.svc.Create(user); err != nil {
		errJSON(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, basicMessage{Message: "user created successfully"})
}

func (uc UserController) get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		errJSON(w, r, ParameterErr{Param: "id", Err: "empty value"})
		return
	}
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		errJSON(w, r, ParameterErr{Param: "id", Err: err.Error()})
		return
	}
	user, err := uc.svc.Get(id)
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, user)
}

func (uc UserController) getAll(w http.ResponseWriter, r *http.Request) {
	user, err := uc.svc.GetAll()
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, user)
}

func (uc UserController) getFiltered(w http.ResponseWriter, r *http.Request) {
	filter := chi.URLParam(r, "filter")
	if filter == "" {
		errJSON(w, r, ParameterErr{Param: "filter", Err: "filter empty"})
		return
	}
	value := chi.URLParam(r, "value")
	if value == "" {
		errJSON(w, r, ParameterErr{Param: "value", Err: "filter value empty"})
		return
	}
	users, err := uc.svc.Find(filter, value)
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, users)
}

func (uc UserController) update(w http.ResponseWriter, r *http.Request) {
	var dto UserUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}
	user := entity.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDay:  dto.BirthDay,
		Username:  dto.Username,
	}
	if err := uc.svc.Update(user); err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, basicMessage{Message: "user updated successfully"})
}

func (uc UserController) delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		errJSON(w, r, ParameterErr{Param: "id", Err: "empty value"})
		return
	}
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		errJSON(w, r, ParameterErr{Param: "id", Err: err.Error()})
		return
	}
	if err := uc.svc.Delete(id); err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, fmt.Sprintf("user %d deleted successfully", id))
}

func (uc UserController) login(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		errJSON(w, r, ParameterErr{Param: "username", Err: "empty value"})
		return
	}
	passwd := chi.URLParam(r, "passwd")
	if passwd == "" {
		errJSON(w, r, ParameterErr{Param: "password", Err: "empty value"})
		return
	}
	user, err := uc.svc.ValidateLogin(username, passwd)
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, UserLoginResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		LastLogin: user.LastLogin.Time,
	})
}
