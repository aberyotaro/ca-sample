package contract

import "github.com/aberyotaro/ca_sample/internal/model"

type User interface {
	Get(id *int64) *model.User
}
