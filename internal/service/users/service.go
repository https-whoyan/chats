package users

import (
	"context"
	"github.com/https-whoyan/chats/internal/domain/entity"
	"github.com/https-whoyan/chats/internal/repository/users"
	"github.com/https-whoyan/chats/internal/usecases/hash"
	"github.com/https-whoyan/chats/internal/usecases/validator"
)

type Service interface {
	Create(ctx context.Context, user *entity.User) error
}

type service struct {
	usersR users.Repository
}

func New(usersR users.Repository) Service {
	return &service{
		usersR: usersR,
	}
}

func (s *service) Create(ctx context.Context, user *entity.User) error {
	v := validator.New().
		Required(user).
		Required(user.Nickname).
		Required(user.Password).
		Between(int(user.Age), 10, 100)
	if err := v.Validate(); err != nil {
		return err
	}
	user.Password = hash.GetHash(user.Password)
	return s.usersR.Create(ctx, user)
}
