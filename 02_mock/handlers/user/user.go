package user

import (
	"context"
	"fmt"

	user "github.com/alexapps/examples/02_mock/storage/user"
)

//go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=UserCreator
type UserCreator interface {
	Create(ctx context.Context, u user.User) (int, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=UserProvider
type UserProvider interface {
	User(ctx context.Context, email string) (*user.User, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=UserNotifier
type UserNotifier interface {
	NotifyUserCreated(ctx context.Context, u user.User) error
}

type Service struct {
	userCreator  UserCreator
	userProvider UserProvider
	userNotifier UserNotifier
}

func (s *Service) CreateUser(ctx context.Context, u user.User) (int, error) {

	// first check if user exists
	foundUser, err := s.userProvider.User(ctx, u.Email)
	if err != nil {
		return 0, fmt.Errorf("can't get user %v: %w", u, err)
	}

	if foundUser != nil {
		return 0, fmt.Errorf("user already present %v", u)
	}

	// create user
	uid, err := s.userCreator.Create(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("can't create user %v: %w", u, err)
	}

	// notify
	if err = s.userNotifier.NotifyUserCreated(ctx, u); err != nil {
		return 0, fmt.Errorf("can't notify user created  %v: %w", u, err)
	}

	return uid, nil
}
