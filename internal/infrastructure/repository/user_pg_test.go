package repository

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepositoryPgTestSuite struct {
	suite.Suite
	// conn *db.PgConn
}

func TestUserRepositoryPgTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryPgTestSuite))
}

func (s *UserRepositoryPgTestSuite) SetupSuite() {
	// cfg := config.NewConfig()
	// conn, err := db.NewPgConn(cfg.Database.Postgres)
	// s.Require().Nil(err)
	// s.conn = conn
}

func (s *UserRepositoryPgTestSuite) TearDownSuite() {
	// s.Assert().NoError(s.conn.Close(),
	// 	"close database connection is required")
}

func (s *UserRepositoryPgTestSuite) TestCreate() {
	// tests := []struct {
	// 	name    string
	// 	user    entity.User
	// 	err     error
	// 	wantErr bool
	// }{
	// 	{
	// 		name:    "Empty",
	// 		user:    entity.User{},
	// 		err:     &EntityEmptyErr{Name: "User"},
	// 		wantErr: true,
	// 	},
	// 	{
	// 		name: "Invalid",
	// 		user: entity.User{
	// 			FirstName: "",
	// 			LastName:  "bar",
	// 			Email:     "foo@example.com",
	// 			Username:  "foouser",
	// 			Passwd:    "foopasswd",
	// 		},
	// 		err:     &InvalidFieldErr{Name: "FirstName", Err: ErrFieldEmpty},
	// 		wantErr: true,
	// 	},
	// 	{
	// 		name: "Created",
	// 		user: entity.User{
	// 			FirstName: "foo",
	// 			LastName:  "bar",
	// 			Email:     "foo@example.com",
	// 			Username:  "foouser",
	// 			Passwd:    "foopasswd",
	// 		},
	// 		err:     nil,
	// 		wantErr: false,
	// 	},
	// }

	// for _, tt := range tests {
	// 	s.Run(tt.name, func() {
	// 		repo := UserRepositoryPg{
	// 			db: s.conn.DB(),
	// 		}
	// 		err := repo.Create(tt.user)
	// 		if tt.wantErr {
	// 			s.Assert().Equal(tt.err, err)
	// 			return
	// 		}
	// 		s.Assert().Nil(err)
	// 	})
	// }
}

func (s *UserRepositoryPgTestSuite) TestRead() {
	// repo := UserRepositoryPg{
	// 	db: s.conn.DB(),
	// }
	// out, err := repo.Read(1)
	// s.Require().Nil(err)

	// s.Assert().Equal(1, out.ID)
	// s.Assert().Equal("foo", out.FirstName)
	// s.Assert().Equal("bar", out.LastName)
	// s.Assert().Equal("foo@example.com", out.Email)
	// s.Assert().Equal("foouser", out.Username)
	// s.Assert().Equal(false, out.Active)
}

func (s *UserRepositoryPgTestSuite) TestReadAll() {
	// repo := UserRepositoryPg{
	// 	db: s.conn.DB(),
	// }
	// out, err := repo.ReadAll()
	// s.Require().Nil(err)
	// s.T().Logf("OUT: %#v", out)
}

func (s *UserRepositoryPgTestSuite) TestUpdate() {
	// repo := UserRepositoryPg{
	// 	db: s.conn.DB(),
	// }
	// err := repo.Update(entity.User{
	// 	ID:        1,
	// 	FirstName: "baz",
	// 	LastName:  "baz",
	// 	Email:     "baz@example.com",
	// 	Username:  "bazuser",
	// 	Passwd:    "bazpasswd",
	// })
	// s.Require().Nil(err)
}

func (s *UserRepositoryPgTestSuite) TestDelete() {
	// repo := UserRepositoryPg{
	// 	db: s.conn.DB(),
	// }
	// err := repo.Delete(1)
	// s.Assert().Nil(err)
}

func TestNewUserRepositoryPg(t *testing.T) {
	// cfg := config.NewConfig().Database.Postgres
	// conn, err := db.NewPgConn(cfg)
	// require.Nil(t, err)
	// defer conn.Close()
	// assert.NotNil(t, conn)
}
