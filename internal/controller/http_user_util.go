package controller

import (
	"fmt"

	"github.com/wizeline/CA-Microservices-Go/internal/service"
)

func parseUserResponse(user service.UserResponse) userResponse {
	return userResponse{
		ID:        fmt.Sprintf("%d", user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		BirthDay:  user.BirthDay.Format(dateFormat),
		Username:  user.Username,
	}
}
