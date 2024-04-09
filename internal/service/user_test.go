package service

import "github.com/wizeline/CA-Microservices-Go/internal/service/mocks"

// We ensure the UserRepo mock object satisfies the UserRepo signature.
var _ UserRepo = &mocks.UserRepo{}
