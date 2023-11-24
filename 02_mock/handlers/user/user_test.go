package user

import (
	"context"
	"testing"

	"github.com/alexapps/examples/02_mock/handlers/user/mocks"
	user "github.com/alexapps/examples/02_mock/storage/user"
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
				ctx: context.TODO(),
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
			// mock.Anything - any values as ctx value for instance
			userProvider.
				On("User", tt.args.ctx, tt.args.u.Email).
				Once().
				Return(nil, nil)

			userCreator.
				On("Create", tt.args.ctx, tt.args.u).
				Once().
				Return(0, nil)

			userNotifier.
				On("NotifyUserCreated", tt.args.ctx, tt.args.u).
				Once().
				Return(nil)

			s := Service{
				userCreator:  userCreator,
				userProvider: userProvider,
				userNotifier: userNotifier,
			}
			_, err := s.CreateUser(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
