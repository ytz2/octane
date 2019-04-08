package controller

import (
	"context"
	"octane/entities"
	"octane/record"
	"octane/repository"
)

// Lotto ...
type Lotto interface {
	GetUser(ctx context.Context, id string) (*record.User, error)
	GetUsers(ctx context.Context, request *entities.GetUsersRequest) ([]record.User, error)
	UpdateUser(ctx context.Context, user *record.User) error
	DeleteUser(ctx context.Context, id string) error
	AddUser(ctx context.Context, user *record.User) error

	GetCategoriesByUserID(ctx context.Context, userId string) ([]record.Category, error)
	GetCategory(ctx context.Context, categoryId string) (*record.Category, error)
	UpdateCategory(ctx context.Context, category *record.Category) error
	DeleteCategory(ctx context.Context, categoryId string) error
	AddCategory(ctx context.Context, category *record.Category) error

	GetContextsByCategoryID(ctx context.Context, categoryID string) ([]record.Context, error)
	GetContext(ctx context.Context, contextId string) (*record.Context, error)
	UpdateContext(ctx context.Context, context *record.Context) error
	DeleteContext(ctx context.Context, ctxId string) error
	AddContext(ctx context.Context, context *record.Context) error
}

type lotto struct {
	r repository.LottoMySQL
}

// NewLotto ...
func NewLotto(l repository.LottoMySQL) Lotto {
	return &lotto{
		r: l,
	}
}
