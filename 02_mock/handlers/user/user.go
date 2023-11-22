package user

import (
	"context"
	"fmt"
)

type User struct {
	Name     string
	Position string
	Email    string
}

//go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=UserCreator
type UserCreator interface {
	Create(ctx context.Context, u User) (int64, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=UserProvider
type UserProvider interface {
	User(ctx context.Context, email string) (*User, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=UserNotifier
type UserNotifier interface {
	NotifyUserCreated(ctx context.Context, u User) error
}

type Service struct {
	UserCreator  UserCreator
	UserProvider UserProvider
	UserNotifier UserNotifier
}

func (s *Service) CreateUser(ctx context.Context, u User) (int64, error) {

	// first check if user exists
	foundUser, err := s.UserProvider.User(ctx, u.Email)
	if err != nil {
		return 0, fmt.Errorf("can't get user %v: %w", u, err)
	}

	if foundUser != nil {
		return 0, fmt.Errorf("user already present %v", u)
	}

	// create user
	uid, err := s.UserCreator.Create(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("can't create user %v: %w", u, err)
	}

	// notify
	if err := s.UserNotifier.NotifyUserCreated(ctx, u); err != nil {
		return 0, fmt.Errorf("can't notify user created  %v: %w", u, err)
	}

	return uid, nil
}
