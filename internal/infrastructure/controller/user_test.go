package controller

// import (
// 	"testing"

// 	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
// 	"github.com/wizeline/CA-Microservices-Go/internal/domain/service/mocks"
// )

// func TestUserCtrl_add(t *testing.T) {
// 	type svc struct {
// 		args entity.User
// 		err  error
// 	}
// 	type req struct {
// 		payload string
// 	}
// 	type resp struct {
// 		code int
// 		body string
// 	}
// 	tests := []struct {
// 		name    string
// 		svc     svc
// 		req     req
// 		resp    resp
// 		err     errHTTP
// 		wantErr bool
// 	}{
// 		// {
// 		// 	name: "Empty",
// 		// 	svc: svc{
// 		// 		err: &repository.EntityEmptyErr{Name: "User"},
// 		// 	},
// 		// 	wantErr: true,
// 		// },
// 		{
// 			name: "Created",
// 			svc: svc{
// 				args: entity.User{
// 					FirstName: "foo",
// 					LastName:  "baz",
// 					Email:     "foo@example.com",
// 					Username:  "foouser",
// 					Passwd:    "foopasswd",
// 				},
// 				err: nil,
// 			},
// 			req: req{
// 				payload: `{"first_name": "foo","last_name": "baz","email": "foo@example.com", "username": "foouser","password": "foopasswd"}`,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mock := mocks.NewUserService(t)
// 			mock.On("Add", tt.svc.dto).Return(tt.svc.err)
// 			// ctrl := NewUserController(mock)

// 			// ts := httptest.NewServer(http.HandlerFunc(ctrl.add))
// 			// defer ts.Close()

// 			// res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer([]byte(tt.req.payload)))
// 			// require.Nil(t, err)

// 			// t.Logf("Code: %v", res.StatusCode)
// 			// t.Logf("Body: %s", res.Body)

// 			// res.Body.Close()
// 			// if err != nil {
// 			// 	t.Fatal(err)
// 			// }

// 			// req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(tt.req.payload)))
// 			// rec := httptest.NewRecorder()
// 			// r := chi.NewRouter()
// 			// r.Post("/users", ctrl.add)
// 			// r.ServeHTTP(rec, req)

// 			// t.Logf("Code: %v", rec.Code)
// 			// t.Logf("Body: %s", rec.Body)

// 			// if tt.wantErr {
// 			// require.NotNil(t, err)
// 			// assert.ErrorIs(t, err, tt.err)
// 			// return
// 			// }
// 			// assert.Equal(t, tt.resp.code, ts.Code)
// 			// assert.Equal(t, tt.resp.body, rec.Body.String())
// 		})
// 	}
// }

// // func TestUserCtrl_get(t *testing.T) {
// // 	type fields struct {
// // 		svc svc.UserService
// // 	}
// // 	type args struct {
// // 		c echo.Context
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		fields  fields
// // 		args    args
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			uc := UserCtrl{
// // 				svc: tt.fields.svc,
// // 			}
// // 			if err := uc.get(tt.args.c); (err != nil) != tt.wantErr {
// // 				t.Errorf("UserCtrl.get() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }

// // func TestUserCtrl_getFiltered(t *testing.T) {
// // 	type fields struct {
// // 		svc svc.UserService
// // 	}
// // 	type args struct {
// // 		c echo.Context
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		fields  fields
// // 		args    args
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			uc := UserCtrl{
// // 				svc: tt.fields.svc,
// // 			}
// // 			if err := uc.getFiltered(tt.args.c); (err != nil) != tt.wantErr {
// // 				t.Errorf("UserCtrl.getFiltered() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }

// // func TestUserCtrl_update(t *testing.T) {
// // 	type fields struct {
// // 		svc svc.UserService
// // 	}
// // 	type args struct {
// // 		c echo.Context
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		fields  fields
// // 		args    args
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			uc := UserCtrl{
// // 				svc: tt.fields.svc,
// // 			}
// // 			if err := uc.update(tt.args.c); (err != nil) != tt.wantErr {
// // 				t.Errorf("UserCtrl.update() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }

// // func TestUserCtrl_delete(t *testing.T) {
// // 	type fields struct {
// // 		svc svc.UserService
// // 	}
// // 	type args struct {
// // 		c echo.Context
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		fields  fields
// // 		args    args
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			uc := UserCtrl{
// // 				svc: tt.fields.svc,
// // 			}
// // 			if err := uc.delete(tt.args.c); (err != nil) != tt.wantErr {
// // 				t.Errorf("UserCtrl.delete() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }

// // func TestUserCtrl_login(t *testing.T) {
// // 	type fields struct {
// // 		svc svc.UserService
// // 	}
// // 	type args struct {
// // 		c echo.Context
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		fields  fields
// // 		args    args
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			uc := UserCtrl{
// // 				svc: tt.fields.svc,
// // 			}
// // 			if err := uc.login(tt.args.c); (err != nil) != tt.wantErr {
// // 				t.Errorf("UserCtrl.login() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }
