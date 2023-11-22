package user_test

import (
	"context"
	"testing"

	"github.com/alexapps/examples/02_mock/handlers/user"
	"github.com/alexapps/examples/02_mock/handlers/user/mocks"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateUser(t *testing.T) {

	type args struct {
		ctx context.Context
		u   user.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base test",
			args: args{
				ctx: context.Background(),
				u: user.User{
					Email: "aaa@mail.eu",
					Name:  "Carlos",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cerate mocks for dependings
			userProvider := mocks.NewUserProvider(t)
			userCreator := mocks.NewUserCreator(t)
			userNotifier := mocks.NewUserNotifier(t)

			// userProvider.On("User", tt.args.ctx, tt.args.u.Email).Return(nil, nil)
			// mock.Anything - any values
			userProvider.On("User", mock.Anything, tt.args.u.Email).Return(nil, nil)
			userCreator.On("Create", tt.args.ctx, tt.args.u).Return(0, nil)
			userNotifier.On("NotifyUserCreated", tt.args.ctx, tt.args.u).Return(nil)

			s := &user.Service{
				UserCreator:  userCreator,
				UserProvider: userProvider,
				UserNotifier: userNotifier,
			}
			_, err := s.CreateUser(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
