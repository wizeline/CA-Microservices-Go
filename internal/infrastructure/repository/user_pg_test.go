package repository

import (
	"testing"

	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/config"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryPgTestSuite struct {
	suite.Suite
	conn *db.PgConn
}

func TestUserRepositoryPgTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryPgTestSuite))
}

func (s *UserRepositoryPgTestSuite) SetupSuite() {
	cfg := config.NewConfig()
	conn, err := db.NewPgConn(cfg.Database.Postgres)
	require.Nil(s.T(), err)
	s.conn = conn
}

func (s *UserRepositoryPgTestSuite) TearDownSuite() {
	assert.NoError(s.T(), s.conn.Close(),
		"close database connection is required")
}

func (s *UserRepositoryPgTestSuite) TestCreate() {
	tests := []struct {
		name    string
		user    entity.User
		err     error
		wantErr bool
	}{
		{
			name:    "Empty",
			user:    entity.User{},
			err:     ErrUserEmpty,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			repo := UserRepositoryPg{
				db: s.conn.DB(),
			}
			err := repo.Create(tt.user)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func (s *UserRepositoryPgTestSuite) TestRead() {

}

func (s *UserRepositoryPgTestSuite) TestReadAll() {

}

func (s *UserRepositoryPgTestSuite) TestUpdate() {

}

func (s *UserRepositoryPgTestSuite) TestDelete() {

}

func TestNewUserRepositoryPg(t *testing.T) {
	cfg := config.NewConfig().Database.Postgres
	conn, err := db.NewPgConn(cfg)
	require.Nil(t, err)
	defer conn.Close()
	assert.NotNil(t, conn)
}
