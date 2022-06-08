package main

import (
	"fmt"
	"github.com/aberyotaro/ca_sample/internal/handler"
	"github.com/aberyotaro/ca_sample/internal/infrastructure"
	"github.com/aberyotaro/ca_sample/internal/use_case"
)

func main() {
	i := infrastructure.NewUserRepository()
	uc := use_case.NewUser(i)
	h := handler.NewUserHandler(uc)

	id := int64(1)
	fmt.Println(h.Handle(&id))
}
