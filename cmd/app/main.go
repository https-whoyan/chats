package main

import (
	"context"
	"github.com/https-whoyan/chats/internal/app"
)

func main() {
	ctx := context.Background()

	goApp := app.NewApp()
	if goApp == nil {
		panic("nil app")
	}
	err := goApp.InitApp(ctx)
	if err != nil {
		panic(err)
	}
	err = goApp.Start()
	if err != nil {
		panic(err)
	}
}
