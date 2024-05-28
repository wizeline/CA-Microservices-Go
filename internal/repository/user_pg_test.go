package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/config"
	"github.com/wizeline/CA-Microservices-Go/internal/db"
	"github.com/wizeline/CA-Microservices-Go/internal/db/migration"
	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/util"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var testValidUser = UserCreateArgs{
	FirstName: "foo",
	LastName:  "bar",
	Birthday:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	Email:     "foo@example.com",
	Username:  "foouser",
	Passwd:    "foopasswd",
}

type UserRepositoryPgTestSuite struct {
	suite.Suite
	conn *db.PgConn
}

func TestUserRepositoryPgTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryPgTestSuite))
}

func (s *UserRepositoryPgTestSuite) SetupSuite() {
	dbCfg := config.NewConfig().Database.Postgres
	dbCfg.Host = "localhost"

	conn, err := db.NewPgConn(dbCfg)
	s.Require().Nil(err)
	s.conn = conn

	migration := migration.CreateUsersTable
	filePath := filepath.Join("./../../migrations/v1", migration.Filename)
	sqlContent, err := os.ReadFile(filePath)
	s.Require().Nil(err)
	s.NoError(migration.Down(conn.DB(), ""))
	s.NoError(migration.Up(conn.DB(), string(sqlContent)))
}

func (s *UserRepositoryPgTestSuite) TearDownSuite() {
	s.Assert().NoError(s.conn.Close(),
		"close database connection is required")
}

func (s *UserRepositoryPgTestSuite) TearDownTest() {
	_, err := s.conn.DB().Exec("TRUNCATE TABLE users RESTART IDENTITY")
	s.Require().Nil(err)

	fmt.Println("==> TEST DATA: removed")
}

func (s *UserRepositoryPgTestSuite) TestCreate() {
	tests := []struct {
		name    string
		args    UserCreateArgs
		wantErr error
	}{
		{
			name:    "Empty",
			args:    UserCreateArgs{},
			wantErr: &EntityEmptyErr{Name: "UserCreateArgs"},
		},
		{
			name: "Invalid",
			args: UserCreateArgs{
				FirstName: "",
				LastName:  "bar",
				Birthday:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				Email:     "foo@example.com",
				Username:  "foouser",
				Passwd:    "foopasswd",
			},
			wantErr: &InvalidFieldErr{Name: "FirstName", Err: util.ErrEmptyValue},
		},
		{
			name:    "Created",
			args:    testValidUser,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := UserRepositoryPg{db: s.conn.DB()}.Create(tt.args)
			if tt.wantErr != nil {
				s.Assert().Equal(err, tt.wantErr)
				return
			}
			s.Assert().Nil(err)
		})
	}
}

func (s *UserRepositoryPgTestSuite) TestRead() {
	repo := UserRepositoryPg{db: s.conn.DB()}
	s.Assert().NoError(repo.Create(testValidUser))

	id := uint64(1)
	out, err := repo.Read(id)
	s.Require().Nil(err)

	s.Assert().Equal(id, out.ID)
	s.Assert().Equal(testValidUser.FirstName, out.FirstName)
	s.Assert().Equal(testValidUser.LastName, out.LastName)
	s.Assert().Equal(testValidUser.Email, out.Email)
	s.Assert().Equal(testValidUser.Username, out.Username)
	s.Assert().Equal(testValidUser.Passwd, out.Passwd)
	s.Assert().Equal(false, out.Active)
}

func (s *UserRepositoryPgTestSuite) TestReadAll() {
	repo := UserRepositoryPg{db: s.conn.DB()}
	s.Assert().NoError(repo.Create(testValidUser))

	out, err := repo.ReadAll()
	s.Require().Nil(err)

	s.Assert().Len(out, 1)
	s.Assert().Equal(uint64(1), out[0].ID)
	s.Assert().Equal(testValidUser.FirstName, out[0].FirstName)
	s.Assert().Equal(testValidUser.LastName, out[0].LastName)
	s.Assert().Equal(testValidUser.Email, out[0].Email)
	s.Assert().Equal(testValidUser.Username, out[0].Username)
	s.Assert().Equal(testValidUser.Passwd, out[0].Passwd)
	s.Assert().Equal(false, out[0].Active)
}

func (s *UserRepositoryPgTestSuite) TestUpdate() {
	repo := UserRepositoryPg{db: s.conn.DB()}
	s.Assert().NoError(repo.Create(testValidUser))

	err := repo.Update(entity.User{
		ID:        1,
		FirstName: "baz",
		LastName:  "baz",
		Email:     "baz@example.com",
		Username:  "bazuser",
		Passwd:    "bazpasswd",
	})
	s.Require().Nil(err)
}

func (s *UserRepositoryPgTestSuite) TestDelete() {
	repo := UserRepositoryPg{db: s.conn.DB()}
	s.Assert().NoError(repo.Create(testValidUser))

	err := repo.Delete(1)
	s.Assert().Nil(err)
}

func TestNewUserRepositoryPg(t *testing.T) {
	dbCfg := config.NewConfig().Database.Postgres
	dbCfg.Host = "localhost"

	conn, err := db.NewPgConn(dbCfg)
	require.Nil(t, err)
	defer conn.Close()
	assert.NotNil(t, conn)
}
