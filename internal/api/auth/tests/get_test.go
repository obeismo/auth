package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/obeismo/auth/internal/api/auth"
	"github.com/obeismo/auth/internal/converter"
	"github.com/obeismo/auth/internal/model"
	"github.com/obeismo/auth/internal/service"
	serviceMocks "github.com/obeismo/auth/internal/service/mocks"
	desc "github.com/obeismo/auth/pkg/auth/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.GetUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id         = gofakeit.Int64()
		name       = gofakeit.Name()
		email      = gofakeit.Email()
		password   = gofakeit.Word()
		date       = gofakeit.Date()
		serviceErr = errors.New("service error")

		modelUserInfo = model.UserInfo{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: "",
			Role:            desc.Role_USER,
		}

		servUserModel = &model.User{
			ID:        id,
			Info:      modelUserInfo,
			CreatedAt: date,
			UpdatedAt: date,
		}

		userAPI = &desc.User{
			Id:        id,
			Info:      converter.ToUserInfoDescFromService(&modelUserInfo),
			CreatedAt: timestamppb.New(date),
			UpdatedAt: timestamppb.New(date),
		}

		req = &desc.GetUserRequest{
			Id: id,
		}

		res = &desc.GetUserResponse{
			User: userAPI,
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetUserResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success get",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, req.Id).Return(servUserModel, nil)
				return mock
			},
		},
		{
			name: "failed get",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, req.Id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// create mocks for service
			authServiceMock := tt.authServiceMock(mc)

			// create api handler
			api := auth.NewServer(authServiceMock)

			// run create test
			resHandler, err := api.Get(tt.args.ctx, tt.args.req)

			// check result
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
