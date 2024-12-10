package app

import (
	"context"
	"io"

	"github.com/https-whoyan/chats/internal/repository/conn"

	usersR "github.com/https-whoyan/chats/internal/repository/users"
	usersS "github.com/https-whoyan/chats/internal/service/users"

	"github.com/jackc/pgx"
)

type App struct {
	starters   []starter
	closers    []io.Closer
	startErrCh chan error

	sqlConn conn.SQLConnection

	usersR usersR.Repository

	usersS usersS.Service
}

func NewApp() *App {
	return &App{
		startErrCh: make(chan error),
	}
}

type initAppFunc = func(ctx context.Context) error

func (a *App) InitApp(ctx context.Context) error {
	var err error
	for _, fn := range []initAppFunc{
		a.initConn,
		a.initRepositories,
	} {
		err = fn(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConn(ctx context.Context) error {
	var err error
	a.sqlConn, err = conn.NewSQLRepository(ctx, pgx.ConnConfig{
		Host: "localhost",
		User: "yan",
		Port: 5432,
	})
	if err != nil {
		return err
	}
	a.addComponent(a.sqlConn)
	return nil
}

func (a *App) initRepositories(_ context.Context) error {
	a.usersR = usersR.NewRepository(a.sqlConn)
	return nil
}

func (a *App) initServices(_ context.Context) error {
	a.usersS = usersS.New(a.usersR)
	return nil
}
