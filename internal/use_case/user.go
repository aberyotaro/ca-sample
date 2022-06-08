package use_case

import (
	"github.com/aberyotaro/ca_sample/internal/contract"
	"github.com/aberyotaro/ca_sample/internal/model"
)

type User struct {
	user contract.User
}

func NewUser(userContract contract.User) *User {
	return &User{user: userContract}
}

func (u *User) Get(id *int64) *model.User {
	return u.user.Get(id)
}
