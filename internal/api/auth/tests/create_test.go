package tests

import (
	"context"
	"testing"

	"errors"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/obeismo/auth/internal/api/auth"
	"github.com/obeismo/auth/internal/converter"
	"github.com/obeismo/auth/internal/model"
	"github.com/obeismo/auth/internal/service"
	serviceMocks "github.com/obeismo/auth/internal/service/mocks"
	desc "github.com/obeismo/auth/pkg/auth/v1"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.CreateUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id              = gofakeit.Int64()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Word()
		passwordConfirm = gofakeit.Word()
		serviceErr      = errors.New("service error")

		userInfoSuccess = &desc.UserInfo{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role_USER,
		}

		// password and passwordConfirm are not equal
		userInfoFailedCase1 = &desc.UserInfo{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            desc.Role_USER,
		}

		// empty name
		userInfoFailedCase2 = &desc.UserInfo{
			Name:            "",
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role_USER,
		}

		// empty email
		userInfoFailedCase3 = &desc.UserInfo{
			Name:            name,
			Email:           "",
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role_USER,
		}

		// empty password
		userInfoFailedCase4 = &desc.UserInfo{
			Name:            name,
			Email:           email,
			Password:        "",
			PasswordConfirm: password,
			Role:            desc.Role_USER,
		}

		reqSuccess = &desc.CreateUserRequest{
			Info: userInfoSuccess,
		}

		reqFailedCase1 = &desc.CreateUserRequest{
			Info: userInfoFailedCase1,
		}

		reqFailedCase2 = &desc.CreateUserRequest{
			Info: userInfoFailedCase2,
		}

		reqFailedCase3 = &desc.CreateUserRequest{
			Info: userInfoFailedCase3,
		}

		reqFailedCase4 = &desc.CreateUserRequest{
			Info: userInfoFailedCase4,
		}

		userModelSuccess = &model.UserInfo{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role_USER,
		}

		res = &desc.CreateUserResponse{
			Id: id,
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateUserResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success create",
			args: args{
				ctx: ctx,
				req: reqSuccess,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, userModelSuccess).Return(id, nil)
				return mock
			},
		},
		{
			name: "failed create case 1",
			args: args{
				ctx: ctx,
				req: reqFailedCase1,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, converter.ToUserInfoServiceFromDesc(userInfoFailedCase1)).Return(0, serviceErr)
				return mock
			},
		},
		{
			name: "failed create case 2",
			args: args{
				ctx: ctx,
				req: reqFailedCase2,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, converter.ToUserInfoServiceFromDesc(userInfoFailedCase2)).Return(0, serviceErr)
				return mock
			},
		},
		{
			name: "failed create case 3",
			args: args{
				ctx: ctx,
				req: reqFailedCase3,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, converter.ToUserInfoServiceFromDesc(userInfoFailedCase3)).Return(0, serviceErr)
				return mock
			},
		},
		{
			name: "failed create case 4",
			args: args{
				ctx: ctx,
				req: reqFailedCase4,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, converter.ToUserInfoServiceFromDesc(userInfoFailedCase4)).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// create mocks for service
			authServiceMock := tt.authServiceMock(mc)

			// create api handler
			api := auth.NewServer(authServiceMock)

			// run create test
			resHandler, err := api.Create(tt.args.ctx, tt.args.req)

			// check result
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
