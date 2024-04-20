package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/controller/mocks"
	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/repository"
	"github.com/wizeline/CA-Microservices-Go/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// We ensure the UserSvc mock object satisfies the UserSvc signature.
var _ UserSvc = &mocks.UserSvc{}

func TestUserControlller_create(t *testing.T) {
	type svcArgs struct {
		user service.UserCreateArgs
	}
	type svcResp struct {
		err error
	}
	type svc struct {
		args svcArgs
		resp svcResp
	}
	tests := []struct {
		name    string
		svc     svc
		req     httpRequestTest
		resp    httpResponseTest
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Payload empty",
			svc:  svc{},
			req: httpRequestTest{
				payload: []byte(""),
			},
			resp: httpResponseTest{
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
			req: httpRequestTest{
				payload: []byte(`{"first_name": "foo","last_name": "baz","username": "foouser"`),
			},
			resp: httpResponseTest{
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
				args: svcArgs{
					user: service.UserCreateArgs{
						FirstName: "foo",
						LastName:  "baz",
						Email:     "foo@example.com",
						BirthDay:  time.Date(1990, time.December, 5, 0, 0, 0, 0, time.UTC),
						Username:  "foouser",
						Passwd:    "foopasswd",
					},
				},
				resp: svcResp{
					err: nil,
				},
			},
			req: httpRequestTest{
				payload: []byte(`{"first_name": "foo","last_name": "baz","email": "foo@example.com", "birthday":"1990-12-05", "username": "foouser","password": "foopasswd"}`),
			},
			resp: httpResponseTest{
				code: http.StatusCreated,
				body: "{\"message\":\"user created successfully\"}\n",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mocks.UserSvc{}
			mock.On("Create", tt.svc.args.user).Return(tt.svc.resp.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(tt.req.payload))
			rec := httptest.NewRecorder()

			ctrl.create(rec, req)

			assert.Equal(t, rec.Code, tt.resp.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.resp.body, rec.Body.String())
		})
	}
}

func TestUserController_get(t *testing.T) {
	type svcArgs struct {
		id uint64
	}
	type svcResp struct {
		user entity.User
		err  error
	}
	type svc struct {
		args svcArgs
		resp svcResp
	}
	tests := []struct {
		name    string
		svc     svc
		req     httpRequestTest
		resp    httpResponseTest
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Empty",
			svc: svc{
				args: svcArgs{id: 0},
				resp: svcResp{
					user: entity.User{},
					err:  nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"id": "",
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParamErrStatus,
				Message: "invalid id parameter: empty value",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			svc: svc{
				args: svcArgs{id: 0},
				resp: svcResp{
					user: entity.User{},
					err:  nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"id": "badid",
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParamErrStatus,
				Message: "invalid id parameter: strconv.ParseUint: parsing \"badid\": invalid syntax",
			},
			wantErr: true,
		},
		{
			name: "Valid",
			svc: svc{
				args: svcArgs{id: 1},
				resp: svcResp{
					user: entity.User{
						ID:        1,
						FirstName: "foo",
						LastName:  "baz",
					},
					err: nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"id": "1",
				},
			},
			resp: httpResponseTest{
				code: http.StatusOK,
				body: "{\"id\":\"1\",\"first_name\":\"foo\",\"last_name\":\"baz\",\"email\":\"\",\"birthday\":\"0001-01-01\",\"username\":\"\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\",\"created_at\":\"0001-01-01 00:00:00 +0000 UTC\",\"updated_at\":\"0001-01-01 00:00:00 +0000 UTC\"}\n",
			},
			err:     errHTTP{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mocks.UserSvc{}
			mock.On("Get", tt.svc.args.id).Return(tt.svc.resp.user, tt.svc.resp.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodGet, "/users?id="+tt.req.params["id"], nil)
			rec := httptest.NewRecorder()

			ctrl.get(rec, req)

			assert.Equal(t, rec.Code, tt.resp.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.resp.body, rec.Body.String())
		})
	}
}

func TestUserController_getAll(t *testing.T) {
	type svcResp struct {
		users []entity.User
		err   error
	}
	type svc struct {
		resp svcResp
	}
	tests := []struct {
		name    string
		svc     svc
		resp    httpResponseTest
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Repository error",
			svc: svc{
				resp: svcResp{
					users: nil,
					err:   &repository.Err{Err: errors.New("some repo error")},
				},
			},
			resp: httpResponseTest{
				code: http.StatusInternalServerError,
			},
			err: errHTTP{
				Code:    http.StatusInternalServerError,
				Status:  repoErrStatus,
				Message: "repository: some repo error",
			},
			wantErr: true,
		},
		{
			name: "Service error",
			svc: svc{
				resp: svcResp{
					users: nil,
					err:   &service.Err{Err: errors.New("some svc error")},
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  svcErrStatus,
				Message: "service: some svc error",
			},
			wantErr: true,
		},
		{
			name: "Valid with no records",
			svc: svc{
				resp: svcResp{
					users: make([]entity.User, 0),
					err:   nil,
				},
			},
			resp: httpResponseTest{
				code: http.StatusOK,
				body: "[]\n",
			},
			err:     errHTTP{},
			wantErr: false,
		},
		{
			name: "Valid with records",
			svc: svc{
				resp: svcResp{
					users: []entity.User{
						{ID: 1, FirstName: "foo"},
						{ID: 2, FirstName: "bar"},
						{ID: 3, FirstName: "baz"},
					},
					err: nil,
				},
			},
			resp: httpResponseTest{
				code: http.StatusOK,
				body: "[{\"id\":\"1\",\"first_name\":\"foo\",\"last_name\":\"\",\"email\":\"\",\"birthday\":\"0001-01-01\",\"username\":\"\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\",\"created_at\":\"0001-01-01 00:00:00 +0000 UTC\",\"updated_at\":\"0001-01-01 00:00:00 +0000 UTC\"},{\"id\":\"2\",\"first_name\":\"bar\",\"last_name\":\"\",\"email\":\"\",\"birthday\":\"0001-01-01\",\"username\":\"\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\",\"created_at\":\"0001-01-01 00:00:00 +0000 UTC\",\"updated_at\":\"0001-01-01 00:00:00 +0000 UTC\"},{\"id\":\"3\",\"first_name\":\"baz\",\"last_name\":\"\",\"email\":\"\",\"birthday\":\"0001-01-01\",\"username\":\"\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\",\"created_at\":\"0001-01-01 00:00:00 +0000 UTC\",\"updated_at\":\"0001-01-01 00:00:00 +0000 UTC\"}]\n",
			},
			err:     errHTTP{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mocks.UserSvc{}
			mock.On("GetAll").Return(tt.svc.resp.users, tt.svc.resp.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			rec := httptest.NewRecorder()

			ctrl.getAll(rec, req)

			assert.Equal(t, rec.Code, tt.resp.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.resp.body, rec.Body.String())
		})
	}
}

func TestUserController_getFiltered(t *testing.T) {
	type svcArgs struct {
		filter string
		value  string
	}
	type svcResp struct {
		users []entity.User
		err   error
	}
	type svc struct {
		args svcArgs
		resp svcResp
	}
	tests := []struct {
		name    string
		svc     svc
		req     httpRequestTest
		resp    httpResponseTest
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Filter empty",
			svc: svc{
				args: svcArgs{filter: "", value: ""},
				resp: svcResp{
					users: nil,
					err:   nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"filter": "",
					"value":  "foo-value",
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParamErrStatus,
				Message: "invalid filter parameter: filter empty",
			},
			wantErr: true,
		},
		{
			name: "Filter value empty",
			svc: svc{
				args: svcArgs{filter: "", value: ""},
				resp: svcResp{
					users: nil,
					err:   nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"filter": "foo-filter",
					"value":  "",
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParamErrStatus,
				Message: "invalid value parameter: filter value empty",
			},
			wantErr: true,
		},
		{
			name: "Service error",
			svc: svc{
				args: svcArgs{filter: "foo-filter", value: "foo-value"},
				resp: svcResp{
					users: nil,
					err:   &service.Err{Err: errors.New("some svc error")},
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"filter": "foo-filter",
					"value":  "foo-value",
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  svcErrStatus,
				Message: "service: some svc error",
			},
			wantErr: true,
		},
		{
			name: "No records",
			svc: svc{
				args: svcArgs{filter: "foo-filter", value: "foo-value"},
				resp: svcResp{
					users: make([]entity.User, 0),
					err:   nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"filter": "foo-filter",
					"value":  "foo-value",
				},
			},
			resp: httpResponseTest{
				code: http.StatusOK,
				body: "[]\n",
			},
			err:     errHTTP{},
			wantErr: false,
		},
		{
			name: "Multiple records",
			svc: svc{
				args: svcArgs{filter: "foo-filter", value: "foo-value"},
				resp: svcResp{
					users: []entity.User{
						{FirstName: "foo"},
						{FirstName: "bar"},
						{FirstName: "baz"},
					},
					err: nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"filter": "foo-filter",
					"value":  "foo-value",
				},
			},
			resp: httpResponseTest{
				code: http.StatusOK,
				body: "[{\"id\":\"0\",\"first_name\":\"foo\",\"last_name\":\"\",\"email\":\"\",\"birthday\":\"0001-01-01\",\"username\":\"\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\",\"created_at\":\"0001-01-01 00:00:00 +0000 UTC\",\"updated_at\":\"0001-01-01 00:00:00 +0000 UTC\"},{\"id\":\"0\",\"first_name\":\"bar\",\"last_name\":\"\",\"email\":\"\",\"birthday\":\"0001-01-01\",\"username\":\"\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\",\"created_at\":\"0001-01-01 00:00:00 +0000 UTC\",\"updated_at\":\"0001-01-01 00:00:00 +0000 UTC\"},{\"id\":\"0\",\"first_name\":\"baz\",\"last_name\":\"\",\"email\":\"\",\"birthday\":\"0001-01-01\",\"username\":\"\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\",\"created_at\":\"0001-01-01 00:00:00 +0000 UTC\",\"updated_at\":\"0001-01-01 00:00:00 +0000 UTC\"}]\n",
			},
			err:     errHTTP{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mocks.UserSvc{}
			mock.On("Find", tt.svc.args.filter, tt.svc.args.value).Return(tt.svc.resp.users, tt.svc.resp.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users?filter=%v&value=%v", tt.req.params["filter"], tt.req.params["value"]), nil)
			rec := httptest.NewRecorder()

			ctrl.getFiltered(rec, req)

			assert.Equal(t, rec.Code, tt.resp.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.resp.body, rec.Body.String())
		})
	}
}

// TODO: Update function for new update ser input
// func TestUserControlller_update(t *testing.T) {
// 	type svcArgs struct {
// 		user entity.User
// 	}
// 	type svcResp struct {
// 		err error
// 	}
// 	type svc struct {
// 		args svcArgs
// 		resp svcResp
// 	}
// 	tests := []struct {
// 		name    string
// 		svc     svc
// 		req     httpRequestTest
// 		resp    httpResponseTest
// 		err     errHTTP
// 		wantErr bool
// 	}{
// 		{
// 			name: "Empty",
// 			svc: svc{
// 				args: svcArgs{
// 					user: entity.User{},
// 				},
// 				resp: svcResp{
// 					err: nil,
// 				},
// 			},
// 			req: httpRequestTest{
// 				payload: []byte(""),
// 			},
// 			resp: httpResponseTest{
// 				code: http.StatusUnsupportedMediaType,
// 			},
// 			err: errHTTP{
// 				Code:    http.StatusUnsupportedMediaType,
// 				Status:  ctrlPayloadErrStatus,
// 				Message: "invalid payload: EOF",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Bad JSON",
// 			req: httpRequestTest{
// 				payload: []byte(`{"id": "123", "first_name": "foo","last_name": "baz","username": "foouser"`),
// 			},
// 			resp: httpResponseTest{
// 				code: http.StatusUnsupportedMediaType,
// 			},
// 			err: errHTTP{
// 				Code:    http.StatusUnsupportedMediaType,
// 				Status:  ctrlPayloadErrStatus,
// 				Message: "invalid payload: unexpected EOF",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Bad ID",
// 			req: httpRequestTest{
// 				payload: []byte(`{"id": "badid", "first_name": "foo","last_name": "baz", "birthday": "1990-12-05", "username": "foouser"}`),
// 			},
// 			resp: httpResponseTest{
// 				code: http.StatusUnsupportedMediaType,
// 			},
// 			err: errHTTP{
// 				Code:    http.StatusUnsupportedMediaType,
// 				Status:  ctrlPayloadErrStatus,
// 				Message: "invalid payload: strconv.ParseUint: parsing \"badid\": invalid syntax",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Updated",
// 			svc: svc{
// 				args: svcArgs{
// 					user: entity.User{
// 						ID:        123,
// 						FirstName: "foo",
// 						LastName:  "baz",
// 						BirthDay:  time.Date(1990, time.December, 5, 0, 0, 0, 0, time.UTC),
// 						Username:  "foouser",
// 					},
// 				},
// 				resp: svcResp{
// 					err: nil,
// 				},
// 			},
// 			req: httpRequestTest{
// 				payload: []byte(`{"id": "123", "first_name": "foo","last_name": "baz", "birthday": "1990-12-05", "username": "foouser"}`),
// 			},
// 			resp: httpResponseTest{
// 				code: http.StatusOK,
// 				body: "{\"message\":\"user 123 updated successfully\"}\n",
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mock := &mocks.UserSvc{}
// 			mock.On("Update", tt.svc.args.user).Return(tt.svc.resp.err)
// 			ctrl := NewUserController(mock)

// 			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(tt.req.payload))
// 			rec := httptest.NewRecorder()

// 			ctrl.update(rec, req)

// 			assert.Equal(t, rec.Code, tt.resp.code)
// 			if tt.wantErr {
// 				var errMsg errHTTP
// 				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
// 				assert.Equal(t, tt.err, errMsg)
// 				return
// 			}
// 			assert.Equal(t, tt.resp.body, rec.Body.String())
// 		})
// 	}
// }

func TestUserControlller_delete(t *testing.T) {
	type svcArgs struct {
		id uint64
	}
	type svcResp struct {
		err error
	}
	type svc struct {
		args svcArgs
		resp svcResp
	}
	tests := []struct {
		name    string
		svc     svc
		req     httpRequestTest
		resp    httpResponseTest
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Empty",
			svc: svc{
				args: svcArgs{
					id: 0,
				},
				resp: svcResp{
					err: nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"id": "",
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParamErrStatus,
				Message: "invalid id parameter: empty value",
			},
			wantErr: true,
		},
		{
			name: "Bad ID",
			svc: svc{
				args: svcArgs{
					id: 0,
				},
				resp: svcResp{
					err: nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"id": "badid",
				},
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  ctrlParamErrStatus,
				Message: "invalid id parameter: strconv.ParseUint: parsing \"badid\": invalid syntax",
			},
			wantErr: true,
		},
		{
			name: "Deleted",
			svc: svc{
				args: svcArgs{
					id: 123,
				},
				resp: svcResp{
					err: nil,
				},
			},
			req: httpRequestTest{
				params: map[string]string{
					"id": "123",
				},
			},
			resp: httpResponseTest{
				code: http.StatusOK,
				body: "{\"message\":\"user 123 deleted successfully\"}\n",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mocks.UserSvc{}
			mock.On("Delete", tt.svc.args.id).Return(tt.svc.resp.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users?id=%v", tt.req.params["id"]), nil)
			rec := httptest.NewRecorder()

			ctrl.delete(rec, req)

			assert.Equal(t, rec.Code, tt.resp.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.resp.body, rec.Body.String())
		})
	}
}

func TestUserControlller_login(t *testing.T) {
	type svcArgs struct {
		username string
		passwd   string
	}
	type svcResp struct {
		user entity.User
		err  error
	}
	type svc struct {
		args svcArgs
		resp svcResp
	}
	tests := []struct {
		name    string
		svc     svc
		req     httpRequestTest
		resp    httpResponseTest
		err     errHTTP
		wantErr bool
	}{
		{
			name: "Payload Empty",
			svc:  svc{},
			req: httpRequestTest{
				payload: []byte(""),
			},
			resp: httpResponseTest{
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
			req: httpRequestTest{
				payload: []byte(`{"username": "foo","password": "some-password"`),
			},
			resp: httpResponseTest{
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
			name: "Repository Error",
			svc: svc{
				args: svcArgs{
					username: "foo",
					passwd:   "some-password",
				},
				resp: svcResp{
					user: entity.User{},
					err:  &repository.Err{Err: errors.New("some repository error")},
				},
			},
			req: httpRequestTest{
				payload: []byte(`{"username": "foo","password": "some-password"}`),
			},
			resp: httpResponseTest{
				code: http.StatusInternalServerError,
			},
			err: errHTTP{
				Code:    http.StatusInternalServerError,
				Status:  repoErrStatus,
				Message: "repository: some repository error",
			},
			wantErr: true,
		},
		{
			name: "Service Error",
			svc: svc{
				args: svcArgs{
					username: "foo",
					passwd:   "some-password",
				},
				resp: svcResp{
					user: entity.User{},
					err:  &service.Err{Err: errors.New("some service error")},
				},
			},
			req: httpRequestTest{
				payload: []byte(`{"username": "foo","password": "some-password"}`),
			},
			resp: httpResponseTest{
				code: http.StatusBadRequest,
			},
			err: errHTTP{
				Code:    http.StatusBadRequest,
				Status:  svcErrStatus,
				Message: "service: some service error",
			},
			wantErr: true,
		},
		{
			name: "Valid",
			svc: svc{
				args: svcArgs{
					username: "foouser",
					passwd:   "foopasswd",
				},
				resp: svcResp{
					user: entity.User{
						FirstName: "foo",
						LastName:  "baz",
						Email:     "foo@example.com",
						BirthDay:  time.Date(1990, time.December, 5, 0, 0, 0, 0, time.UTC),
						Username:  "foouser",
						Passwd:    "foopasswd",
					},
					err: nil,
				},
			},
			req: httpRequestTest{
				payload: []byte(`{"username": "foouser","password": "foopasswd"}`),
			},
			resp: httpResponseTest{
				code: http.StatusOK,
				body: "{\"id\":\"0\",\"first_name\":\"foo\",\"last_name\":\"baz\",\"email\":\"foo@example.com\",\"username\":\"foouser\",\"last_login\":\"0001-01-01 00:00:00 +0000 UTC\"}\n",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mocks.UserSvc{}
			mock.On("ValidateLogin", tt.svc.args.username, tt.svc.args.passwd).Return(tt.svc.resp.user, tt.svc.resp.err)
			ctrl := NewUserController(mock)

			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(tt.req.payload))
			rec := httptest.NewRecorder()

			ctrl.login(rec, req)

			assert.Equal(t, rec.Code, tt.resp.code)
			if tt.wantErr {
				var errMsg errHTTP
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &errMsg))
				assert.Equal(t, tt.err, errMsg)
				return
			}
			assert.Equal(t, tt.resp.body, rec.Body.String())
		})
	}
}
