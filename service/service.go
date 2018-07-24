package service

import (
	"github.com/LuanNgoMinh/example-go/service/category"
	"github.com/LuanNgoMinh/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
}
