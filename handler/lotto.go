package handler

import (
	"context"
	"errors"
	"octane/controller"
	lpb "octane/grpc/lotto"
	"octane/record"
)

var limitError = errors.New("RateLimit exceeded")

// Lotto ...
type Lotto interface {
	// User related
	GetUser(ctx context.Context, in *lpb.GetUserRequest) (*lpb.User, error)
	UpdateUser(ctx context.Context, in *lpb.UpdateUserRequest) (*lpb.Void, error)
	AddUser(ctx context.Context, in *lpb.AddUserRequest) (*lpb.Void, error)
	DeleteUser(ctx context.Context, in *lpb.DeleteUserRequest) (*lpb.Void, error)
	GetUsers(ctx context.Context, in *lpb.GetUsersRequest) (*lpb.GetUserResponse, error)

	// Category related
	GetCategoriesByUserID(ctx context.Context, in *lpb.GetCategoriesByUserIDRequest) (*lpb.Categories, error)
	GetCategory(ctx context.Context, in *lpb.GetCategoryRequest) (*lpb.Category, error)
	UpdateCategory(ctx context.Context, in *lpb.UpdateCategoryRequest) (*lpb.Void, error)
	DeleteCategory(ctx context.Context, in *lpb.DeleteCategoryRequest) (*lpb.Void, error)
	AddCategory(ctx context.Context, in *lpb.AddCategoryRequest) (*lpb.Void, error)

	// Context related
	GetContextsByCategoryID(ctx context.Context, in *lpb.GetContextsByCategoryIDRequest) (*lpb.Contexts, error)
	GetContext(ctx context.Context, in *lpb.GetContextRequest) (*lpb.Context, error)
	UpdateContext(ctx context.Context, in *lpb.UpdateContextRequest) (*lpb.Void, error)
	DeleteContext(ctx context.Context, in *lpb.DeleteContextRequest) (*lpb.Void, error)
	AddContext(ctx context.Context, in *lpb.AddContextRequest) (*lpb.Void, error)
}

type lotto struct {
	c controller.Lotto
	r RateLimiter
}

// NewLotto ...
func NewLotto(l controller.Lotto, config *record.Config) Lotto {
	return &lotto{
		c: l,
		r: NewRateLimiter(config),
	}
}
