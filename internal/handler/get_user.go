package handler

import (
	"github.com/aberyotaro/ca_sample/internal/model"
	"github.com/aberyotaro/ca_sample/internal/use_case"
)

type User struct {
	useCase *use_case.User
}

func NewUserHandler(useCase *use_case.User) *User {
	return &User{useCase: useCase}
}

func (u *User) Handle(id *int64) *model.User {
	return u.useCase.Get(id)
}
