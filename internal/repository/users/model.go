package users

import "github.com/https-whoyan/chats/internal/domain/entity"

type model struct {
	Nickname       string `db:"nickname"`
	Age            uint   `db:"age"`
	HashedPassword string `db:"hashed_pass"`
}

func (m model) convert() *entity.User {
	return &entity.User{
		Nickname: m.Nickname,
		Age:      m.Age,
		Password: m.HashedPassword,
	}
}

type models []model

func (mm models) convert() []*entity.User {
	users := make([]*entity.User, len(mm))
	for i, mUser := range mm {
		users[i] = mUser.convert()
	}
	return users
}
