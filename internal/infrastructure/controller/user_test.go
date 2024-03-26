package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/domain/service/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserControlller_create(t *testing.T) {
	type svc struct {
		args entity.User
		err  error
	}
	type req struct {
		payload string
	}
	type res struct {
		code int
		body string
	}
	tests := []struct {
		name    string
		svc     svc
		req     req
		res     res
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Empty",
			svc: svc{
				args: entity.User{},
				err:  nil,
			},
			req: req{
				payload: "",
			},
			res: res{
				code: http.StatusUnsupportedMediaType,
			},
			err: errHTTP{
				Code:    http.StatusUnsupportedMediaType,
				Status:  ctrlPayloadErrStatus,
				Message: "invalid payload: EOF",
			},
			wantErr: true,
		},
		{
			name: "Bad JSON",
			req: req{
				payload: `{"first_name": "foo","last_name": "baz","username": "foouser"`,
			},
			res: res{
				code: http.StatusUnsupportedMediaType,
			},
			err: errHTTP{
				Code:    http.StatusUnsupportedMediaType,
				Status:  ctrlPayloadErrStatus,
				Message: "invalid payload: unexpected EOF",
			},
			wantErr: true,
		},
		{
			name: "Created",
			svc: svc{
				args: entity.User{
					FirstName: "foo",
					LastName:  "baz",
					Email:     "foo@example.com",
					Username:  "foouser",
					Passwd:    "foopasswd",
				},
				err: nil,
			},
			req: req{
				payload: `{"first_name": "foo","last_name": "baz","email": "foo@example.com", "username": "foouser","password": "foopasswd"}`,
			},
			res: res{
				code: http.StatusCreated,
				body: "{\"message\":\"user created successfully\"}\n",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mocks.UserService{}
			mock.On("Create", tt.svc.args).Return(tt.svc.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(tt.req.payload)))
			rec := httptest.NewRecorder()

			ctrl.create(rec, req)

			assert.Equal(t, rec.Code, tt.res.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.res.body, rec.Body.String())
		})
	}
}

func TestUserController_get(t *testing.T) {
	type svcArgs struct {
		id uint64
	}
	type svcRes struct {
		user entity.User
		err  error
	}
	type svc struct {
		args svcArgs
		res  svcRes
	}
	type param struct {
		id string
	}
	type res struct {
		code int
		body string
	}
	tests := []struct {
		name    string
		svc     svc
		param   param
		res     res
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Empty",
			svc: svc{
				args: svcArgs{id: 0},
				res: svcRes{
					user: entity.User{},
					err:  nil,
				},
			},
			param: param{
				id: "",
			},
			res: res{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParameterErrStatus,
				Message: "invalid id parameter: empty value",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			svc: svc{
				args: svcArgs{id: 0},
				res: svcRes{
					user: entity.User{},
					err:  nil,
				},
			},
			param: param{
				id: "badid",
			},
			res: res{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParameterErrStatus,
				Message: "invalid id parameter: strconv.ParseUint: parsing \"badid\": invalid syntax",
			},
			wantErr: true,
		},
		{
			name: "Valid",
			svc: svc{
				args: svcArgs{id: 1},
				res: svcRes{
					user: entity.User{
						ID:        1,
						FirstName: "foo",
						LastName:  "baz",
					},
					err: nil,
				},
			},
			param: param{
				id: "1",
			},
			res: res{
				code: http.StatusOK,
				body: "{\"ID\":1,\"FirstName\":\"foo\",\"LastName\":\"baz\",\"Email\":\"\",\"BirthDay\":\"0001-01-01T00:00:00Z\",\"Username\":\"\",\"Passwd\":\"\",\"Active\":false,\"LastLogin\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false}}\n",
			},
			err:     errHTTP{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock := mocks.NewUserService(t)
			mock := &mocks.UserService{}
			mock.On("Get", tt.svc.args.id).Return(tt.svc.res.user, tt.svc.res.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodGet, "/users?id="+tt.param.id, nil)
			rec := httptest.NewRecorder()

			ctrl.get(rec, req)

			assert.Equal(t, rec.Code, tt.res.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.res.body, rec.Body.String())
		})
	}
}
