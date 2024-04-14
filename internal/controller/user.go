package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/entity"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var _ HTTP = &UserController{}

// userCreateReq represents the data transfer object requested for creating a user.
type userCreateReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	BirthDay  string `json:"birthday"`
	Username  string `json:"username"`
	Passwd    string `json:"password"`
}

// userUpdateReq represents the data transfer object requested for updating a user.
type userUpdateReq struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDay  string `json:"birthday"`
	Username  string `json:"username"`
}

type userResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	BirthDay  string `json:"birthday"`
	Username  string `json:"username"`
	LastLogin string `json:"last_login"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// userLoginReq represents the data transfer object requested for login a user.
type userLoginReq struct {
	Username string `json:"username"`
	Passwd   string `json:"password"`
}

// userLoginResponse represents the data transfer object response for a logged user.
type userLoginResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	LastLogin string `json:"last_login"`
}

type UserSvc interface {
	Create(user entity.User) error
	Get(id uint64) (entity.User, error)
	GetAll() ([]entity.User, error)
	Find(filter, value string) ([]entity.User, error)
	Update(user entity.User) error
	Delete(id uint64) error

	Activate(id uint64) error
	ChangeEmail(id uint64, email string) error
	ChangePasswd(id uint64, passwd string) error
	IsActive(id uint64) (bool, error)
	ValidateLogin(username string, passwd string) (entity.User, error)
}

type UserController struct {
	svc UserSvc
}

func NewUserController(svc UserSvc) UserController {
	return UserController{
		svc: svc,
	}
}

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
	var dto userCreateReq
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}
	birthDay, err := time.Parse(dateFormat, dto.BirthDay)
	if err != nil {
		errJSON(w, r, &PayloadErr{err})
		return
	}

	user := entity.User{
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
	render.JSON(w, r, newUserResponse(user))
}

func (uc UserController) getAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.svc.GetAll()
	if err != nil {
		errJSON(w, r, err)
		return
	}

	usersResp := make([]userResponse, 0)
	for _, u := range users {
		usersResp = append(usersResp, newUserResponse(u))
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
		usersResp = append(usersResp, newUserResponse(u))
	}

	render.JSON(w, r, usersResp)
}

func (uc UserController) update(w http.ResponseWriter, r *http.Request) {
	var dto userUpdateReq
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

	user := entity.User{
		ID:        idUint,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDay:  birthDay,
		Username:  dto.Username,
	}
	if err := uc.svc.Update(user); err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, basicMessage{Message: fmt.Sprintf("user %d updated successfully", user.ID)})
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
	var dto userLoginReq
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
		LastLogin: user.LastLogin.Time.String(),
	})
}

func newUserResponse(user entity.User) userResponse {
	return userResponse{
		ID:        fmt.Sprintf("%d", user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		BirthDay:  user.BirthDay.Format(dateFormat),
		Username:  user.Username,
		LastLogin: user.LastLogin.Time.String(),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
	}
}
