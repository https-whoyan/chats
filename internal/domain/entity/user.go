package entity

type User struct {
	Nickname string `json:"nickname"`
	Age      uint   `json:"age"`
	Password string `json:"password"`
}
