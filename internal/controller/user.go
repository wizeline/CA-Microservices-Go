package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// We guarantee that the requirements of the HTTP controller are met
var _ HTTP = &UserController{}

// userCreateRequest represents the data transfer object requested for creating a user
type userCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	BirthDay  string `json:"birthday"`
	Username  string `json:"username"`
	Passwd    string `json:"password"`
}

// userUpdateRequest represents the data transfer object requested for updating a user
type userUpdateRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDay  string `json:"birthday"`
	Username  string `json:"username"`
}

// userResponse represents the data transfer object response for a default user
type userResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	BirthDay  string `json:"birthday"`
	Username  string `json:"username"`
}

// userLoginRequest represents the data transfer object requested for login a user
type userLoginRequest struct {
	Username string `json:"username"`
	Passwd   string `json:"password"`
}

// userLoginResponse represents the data transfer object response for a logged user
type userLoginResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	LastLogin string `json:"last_login"`
}

// UserService is an abstraction of the UserService dependecy used by the UserController
type UserService interface {
	Create(args service.UserCreateArgs) error
	Get(id uint64) (service.UserResponse, error)
	GetAll() ([]service.UserResponse, error)
	Find(filter, value string) ([]entity.User, error)
	Update(args service.UserUpdateArgs) error
	Delete(id uint64) error

	Activate(id uint64) error
	ChangeEmail(id uint64, email string) error
	ChangePasswd(id uint64, passwd string) error
	IsActive(id uint64) (bool, error)
	ValidateLogin(username string, passwd string) (service.UserLoginResponse, error)
}

// UserController is the user controller representation.
type UserController struct {
	svc UserService
}

// NewUserController returns a new UserController implementation.
func NewUserController(svc UserService) UserController {
	return UserController{
		svc: svc,
	}
}

// SetRoutes sets a fresh middleware stack to configure the handle functions of UserController and mounts them to the given subrouter.
func (uc UserController) SetRoutes(r chi.Router) {
	r.Post("/users", uc.create)
	r.Get("/user", uc.get)
	r.Get("/users", uc.getAll)
	r.Get("/users/filter", uc.getFiltered)
	r.Put("/users", uc.update)
	r.Delete("/users", uc.delete)

	r.Post("/login", uc.login)
}

func (uc UserController) create(w http.ResponseWriter, r *http.Request) {
	var dto userCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}
	birthDay, err := time.Parse(dateFormat, dto.BirthDay)
	if err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}

	user := service.UserCreateArgs{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		BirthDay:  birthDay,
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
	id := r.URL.Query().Get("id")
	if id == "" {
		errJSON(w, r, &ParameterErr{Param: "id", Err: "empty value"})
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errJSON(w, r, &ParameterErr{Param: "id", Err: err.Error()})
		return
	}
	user, err := uc.svc.Get(idUint)
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, parseUserResponse(user))
}

func (uc UserController) getAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.svc.GetAll()
	if err != nil {
		errJSON(w, r, err)
		return
	}

	usersResp := make([]userResponse, 0)
	for _, u := range users {
		usersResp = append(usersResp, parseUserResponse(u))
	}

	render.JSON(w, r, usersResp)
}

func (uc UserController) getFiltered(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	if filter == "" {
		errJSON(w, r, &ParameterErr{Param: "filter", Err: "filter empty"})
		return
	}
	value := r.URL.Query().Get("value")
	if value == "" {
		errJSON(w, r, &ParameterErr{Param: "value", Err: "filter value empty"})
		return
	}
	users, err := uc.svc.Find(filter, value)
	if err != nil {
		errJSON(w, r, err)
		return
	}

	usersResp := make([]userResponse, 0)
	for _, u := range users {
		usersResp = append(usersResp, userResponse{
			ID:        fmt.Sprintf("%d", u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			BirthDay:  u.BirthDay.Format(dateFormat),
			Username:  u.Username,
		})
	}

	render.JSON(w, r, usersResp)
}

func (uc UserController) update(w http.ResponseWriter, r *http.Request) {
	var dto userUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}
	idUint, err := strconv.ParseUint(dto.ID, 10, 64)
	if err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}
	birthDay, err := time.Parse(dateFormat, dto.BirthDay)
	if err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}
	userArgs := service.UserUpdateArgs{
		ID:        idUint,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDay:  birthDay,
	}
	if err := uc.svc.Update(userArgs); err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, basicMessage{Message: fmt.Sprintf("user %d updated successfully", userArgs.ID)})
}

func (uc UserController) delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		errJSON(w, r, &ParameterErr{Param: "id", Err: "empty value"})
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errJSON(w, r, &ParameterErr{Param: "id", Err: err.Error()})
		return
	}
	if err := uc.svc.Delete(idUint); err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, basicMessage{Message: fmt.Sprintf("user %d deleted successfully", idUint)})
}

func (uc UserController) login(w http.ResponseWriter, r *http.Request) {
	var dto userLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}

	user, err := uc.svc.ValidateLogin(dto.Username, dto.Passwd)
	if err != nil {
		errJSON(w, r, err)
		return
	}

	render.JSON(w, r, userLoginResponse{
		ID:        fmt.Sprintf("%d", user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		LastLogin: user.LastLogin.String(),
	})
}
