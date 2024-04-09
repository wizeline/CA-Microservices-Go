package service

import "github.com/wizeline/CA-Microservices-Go/internal/service/mocks"

// We ensure the UserRepository mock object satisfies the UserRepository signature.
var _ UserRepository = &mocks.UserRepository{}
