package handler

import svc "github.com/wizeline/CA-Microservices-Go/internal/domain/service"

type UserHandler struct {
	svc svc.UserService
}

func NewUserHandler(svc svc.UserService) UserHandler {
	return UserHandler{
		svc: svc,
	}
}

// TODO: Implement User handlers
