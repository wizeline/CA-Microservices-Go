package repository

import "github.com/wizeline/CA-Microservices-Go/internal/infrastructure/repository/mocks"

var _ PgConn = &mocks.PgConn{}
