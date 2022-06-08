package infrastructure

import (
	"github.com/aberyotaro/ca_sample/internal/contract"
	"github.com/aberyotaro/ca_sample/internal/model"
)

type UserRepository struct{}

func NewUserRepository() contract.User {
	return &UserRepository{}
}

func (u *UserRepository) Get(id *int64) *model.User {
	return &model.User{
		Id:   *id,
		Name: "Taro Yamada",
	}
}
