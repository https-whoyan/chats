package users

import (
	"context"
	"github.com/https-whoyan/chats/internal/domain/entity"
	"github.com/https-whoyan/chats/internal/repository/conn"
)

type Repository interface {
	Create(ctx context.Context, user *entity.User) error
}

type repository struct {
	conn conn.SQLConnection
}

func NewRepository(conn conn.SQLConnection) Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Create(ctx context.Context, user *entity.User) error {
	return r.conn.Exec(ctx, createUserStmt, user.Nickname, user.Age, user.Password)
}
